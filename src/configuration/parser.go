package configuration

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.rbx.com/roblox/entity-schema-generator/enums"
	"github.rbx.com/roblox/entity-schema-generator/flags"
	"github.rbx.com/roblox/entity-schema-generator/helpers"
	"github.rbx.com/roblox/entity-schema-generator/models"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var ErrConfigurationDirectoryNotSpecified = errors.New("configuration directory not specified")

func parseEntityFileDependingOnExtension(fileName string) (*models.Entity, error) {
	if fileName == "" {
		return nil, nil
	}

	fileExtension := path.Ext(fileName)

	switch fileExtension {
	case ".json":
		return parseEntityJSONFile(fileName)
	case ".yaml":
		return parseEntityYAMLFile(fileName)
	case ".yml":
		return parseEntityYAMLFile(fileName)
	default:
		return nil, nil
	}
}

func parseDatabaseFileDependingOnExtension(fileName string) (*models.Database, error) {
	if fileName == "" {
		return nil, nil
	}

	fileExtension := path.Ext(fileName)

	switch fileExtension {
	case ".json":
		return parseDatabaseJSONFile(fileName)
	case ".yaml":
		return parseDatabaseYAMLFile(fileName)
	case ".yml":
		return parseDatabaseYAMLFile(fileName)
	default:
		return nil, nil
	}
}

var (
	ErrNilEntity                                               = errors.New("entity is nil")
	ErrEntityNameNotSpecified                                  = errors.New("entity name not specified")
	ErrDatabaseNotSpecified                                    = errors.New("database not specified")
	ErrTableNotSpecified                                       = errors.New("table not specified")
	ErrIDPropertyNotSpecified                                  = errors.New("id property not specified")
	ErrPropertiesNotSpecified                                  = errors.New("properties not specified")
	ErrIDPropertyNotNumberType                                 = errors.New("id property is not a number type")
	ErrPropertyNameNotSpecified                                = errors.New("property name not specified")
	ErrPassivePropertyNameNotSpecified                         = errors.New("passive property name not specified")
	ErrPassivePropertyValueNotSpecified                        = errors.New("passive property value not specified")
	ErrPassivePropertyCannotBeIDProperty                       = errors.New("passive property cannot be the id property")
	ErrPassivePropertyCannotBeMethodArgument                   = errors.New("passive property cannot be a method argument")
	ErrPassivePropertyIsNotProperty                            = errors.New("passive property is not a property")
	ErrMethodNameNotSpecified                                  = errors.New("method name not specified")
	ErrParameterNameNotSpecified                               = errors.New("parameter name not specified")
	ErrLookupMethodNoParameters                                = errors.New("lookup methods must have at least one parameter")
	ErrInvalidMemcachedGroupName                               = errors.New("invalid memcached group name setting, format must be <ClassName>:<PropertyName>")
	ErrPredefinedGetMethodName                                 = errors.New("predefined get method name not specified")
	ErrPredefinedValueName                                     = errors.New("predefined value name not specified")
	ErrCannotDeterminePredefinedLookupKey                      = errors.New("cannot determine predefined lookup key")
	ErrCannotDeterminePredefinedLookupKeyWithMultipleArguments = errors.New("cannot determine predefined lookup key with multiple arguments")
	ErrBadVersionNumber                                        = errors.New("only entities v1 and v2 are supported")
	ErrCreatedNotFoundOrWrongType                              = errors.New("created property not found or wrong type")
	ErrMustGetDependentMethodNotSpecified                      = errors.New("must-get dependent method not specified")
	ErrPropertyNotVarBinaryWithNoLength                        = errors.New("property must specify a length if it is not a varbinary")
	ErrParameterNotVarBinaryWithNoLength                       = errors.New("parameter must specify a length if it is not a varbinary")

	ValidForeignKeySyntax = regexp.MustCompile(`^(\[([a-zA-Z0-9_]+)\]\.)?\[dbo\]\.\[([a-zA-Z0-9_]+)\]\.\[([a-zA-Z0-9_]+)\]$`)
)

func toPascalCase(str string) string {
	return strings.ToUpper(str[:1]) + str[1:]
}

func toCamelCase(s string) string {
	// Remove all characters that are not alphanumeric or spaces or underscores
	s = regexp.MustCompile("[^a-zA-Z0-9_ ]+").ReplaceAllString(s, "")

	// Replace all underscores with spaces
	s = strings.ReplaceAll(s, "_", " ")

	// Title case s
	s = cases.Title(language.AmericanEnglish, cases.NoLower).String(s)

	// Remove all spaces
	s = strings.ReplaceAll(s, " ", "")

	// Lowercase the first letter
	if len(s) > 0 {
		s = strings.ToLower(s[:1]) + s[1:]
	}

	return s
}

func getStoredProcedureName(methodType enums.MethodType, entityName string, parameters []*models.Parameter) string {
	var storedProcedureName string

	switch methodType {
	case enums.MethodTypeGetCollection:
		if len(parameters) == 0 {
			storedProcedureName = "GetAll" + entityName + "IDs"
		} else {
			storedProcedureName = "Get" + entityName + "IDsBy"
			for _, parameter := range parameters {
				storedProcedureName += toPascalCase(parameter.Name) + "And"
			}

			storedProcedureName = storedProcedureName[:len(storedProcedureName)-3]
		}
	case enums.MethodTypeGetCollectionExclusive:
		return getStoredProcedureName(enums.MethodTypeGetCollection, entityName, parameters)
	case enums.MethodTypeGetCollectionPaged:
		return getStoredProcedureName(enums.MethodTypeGetCollection, entityName, parameters) + "_Paged"
	case enums.MethodTypeLookup:
		storedProcedureName = "Get" + entityName + "By"
		for _, parameter := range parameters {
			storedProcedureName += toPascalCase(parameter.Name) + "And"
		}

		storedProcedureName = storedProcedureName[:len(storedProcedureName)-3]
	case enums.MethodTypeGetOrCreate:
		storedProcedureName = "GetOrCreate" + entityName
	case enums.MethodTypeMultiGet:
		storedProcedureName = "Get" + entityName + "sByIDs"
	case enums.MethodTypeGetCount:
		storedProcedureName = "GetTotalNumberOf" + entityName

		if len(parameters) != 0 {
			storedProcedureName += "By"
			for _, parameter := range parameters {
				storedProcedureName += toPascalCase(parameter.Name) + "And"
			}

			storedProcedureName = storedProcedureName[:len(storedProcedureName)-3]
		}
	}

	return storedProcedureName
}

func determineKeyRaw(startIndex int, parameters []*models.Parameter) (int, string) {
	// lookup keys and state token keys are like this:
	// {paramNamePascalCase}_{formatIndex}_...
	// e.g. GetPlayerIDByPlayerNameAndPlayerID -> PlayerName_{0}_PlayerID_{1}

	// Full string:
	// string.Format("PlayerName_{0}_PlayerID_{1}", PlayerName, PlayerID)

	var key string
	var lastIndex int

	for i, parameter := range parameters {
		index := strconv.Itoa(startIndex + i)

		key += toPascalCase(parameter.Name) + ":{" + index + "}_"
		lastIndex = i + 1
	}

	key = key[:len(key)-1]

	return lastIndex, key
}

func determineKey(startIndex int, parameters []*models.Parameter, pascalize bool) (int, string) {
	// lookup keys and state token keys are like this:
	// {paramNamePascalCase}_{formatIndex}_...
	// e.g. GetPlayerIDByPlayerNameAndPlayerID -> PlayerName_{0}_PlayerID_{1}

	// Full string:
	// string.Format("PlayerName_{0}_PlayerID_{1}", PlayerName, PlayerID)

	var params []string

	for _, parameter := range parameters {
		if pascalize {
			params = append(params, toPascalCase(parameter.Name))
		} else {
			params = append(params, parameter.Name)
		}
	}

	lastIndex, key := determineKeyRaw(startIndex, parameters)

	return lastIndex, fmt.Sprintf("string.Format(\"%s\", %s)", key, strings.Join(params, ", "))
}

func getFirstElement(m map[string]string) (key string, value string) {
	for k, v := range m {
		return k, v
	}

	return "", ""
}

func existsInProperties(properties []*models.Property, name string) (bool, *models.Property) {
	for _, property := range properties {
		if property.Name == name {
			return true, property
		}
	}

	return false, nil
}

func existsInArguments(arguments []*models.Parameter, name string) bool {
	for _, argument := range arguments {
		if toPascalCase(argument.Name) == name {
			return true
		}
	}

	return false
}

func parsePassiveValueBasedOnType(value string, sqlDbType enums.SqlDbType) string {
	switch sqlDbType {
	case enums.SqlDbTypeBit:
		if value == "true" {
			return "1"
		}

		return "0"
	case enums.SqlDbTypeDateTime:
		return fmt.Sprintf("'%s'", value)
	case enums.SqlDbTypeDateTimeOffset:
		return fmt.Sprintf("'%s'", value)
	case enums.SqlDbTypeDecimal:
		return value
	case enums.SqlDbTypeFloat:
		return value
	case enums.SqlDbTypeInt:
		return value
	case enums.SqlDbTypeVarChar:
		return fmt.Sprintf("'%s'", value)
	case enums.SqlDbTypeNVarChar:
		return fmt.Sprintf("N'%s'", value)
	case enums.SqlDbTypeTime:
		return fmt.Sprintf("'%s'", value)
	case enums.SqlDbTypeTinyInt:
		return value
	case enums.SqlDbTypeUniqueIdentifier:
		return fmt.Sprintf("'%s'", value)
	case enums.SqlDbTypeBinary:
		return fmt.Sprintf("0x%s", value)
	case enums.SqlDbTypeVarBinary:
		return fmt.Sprintf("0x%s", value)
	default:
		return fmt.Sprintf("'%s'", value)
	}
}

func validateNewEntity(entity *models.Entity) error {
	var err error
	if err = preValidate(entity); err != nil {
		return err
	}

	if err = validateIdProperty(entity); err != nil {
		return err
	}

	if err = validateRemoteCacheable(entity); err != nil {
		return err
	}

	setInitialLookups(entity)

	// Get the argument name from the method name. (only if it is a GetByXXX method)
	// Get the argument name from the method name.
	// e.g. GetByPlayerID -> PlayerID is the argument name.
	// ignore multiple "And" arguments.
	// Construct the arguments.
	// e.g. ComputerId = GetByValue(ComputerValue)
	// e.g. ComputerId = GetByValueAndType(ComputerValue, ComputerType)
	// get the first key of the map.
	// there should only be one key.
	if err = validatePredefined(entity); err != nil {
		return err
	}

	// Check properties.
	// Parse the foreign key into regex.
	// Construct the foreign key. If the database is not specified, then it will default to the current database.
	// [database].dbo.[table].[column]
	// e.g. [Roblox].dbo.[Users].[UserID]
	// Foreign key constraint key.
	// e.g. FK_AccountIPAddressesV3_IPAddresses_IPAddressID
	// format: FK_{table}_{foreignTable}_{column}
	if err = validateProperties(entity); err != nil {
		return err
	}

	// Check methods.
	// Wacky parsing on stored procedure names.
	// Cases:
	// Collections:
	// If NO parameters, then it is most likely a get-all operation, hence: GetAll{{EntityName}}IDs
	// If there are parameters, then it is most likely a get-by operation, hence: Get{{EntityName}}IDsBy{{Parameter1Name}}And{{Parameter2Name}}And{{Parameter3Name}} etc.
	// Paged collections follow the same rules as collections, but with the word "_Paged" appended to the end.
	// Lookups:
	// Lookups require at least one parameter, hence: Get{{EntityName}}IDBy{{Parameter1Name}} - if there are more parameters, they are appended with "And" in between.
	// GetOrCreate:
	// GetOrCreate does not take into account any parameters, hence: GetOrCreate{{EntityName}}, if you have more than 1 type of GetOrCreate, then manually specify the name.
	// MultiGet:
	// MultiGet just takes a list of IDs, hence: Get{{EntityName}}sByIDs
	// Error if there are no parameters.
	// If the stored procedure is not specified, then we will construct it.
	// no arguments
	if err = validateMethods(entity); err != nil {
		return err
	}

	if entity.CacheabilitySettings == nil {
		entity.CacheabilitySettings = models.DefaultCacheabilitySettings()
	}

	return nil
}

func validateMethods(entity *models.Entity) error {
	for _, method := range entity.Methods {
		if err := prevalidateMethod(method); err != nil {
			return err
		}

		if method.Type == enums.MethodTypeGetCount {
			if err := handleGetCount(method); err != nil {
				return err
			}
		}

		if method.Type == enums.MethodTypeMustGet {
			if method.DependentMethod == "" {
				return ErrMustGetDependentMethodNotSpecified
			}
		}

		applySharedMethodData(method, entity)

		if err := validatePassiveProperties(method, entity); err != nil {
			return err
		}

		if method.Type == enums.MethodTypeLookup {
			if err := handleLookup(method, entity); err != nil {
				return err
			}
		}

		if method.Type == enums.MethodTypeGetOrCreate {
			handleGetOrCreate(method, entity)
		}

		if method.StoredProcedure == "" {
			method.StoredProcedure = getStoredProcedureName(method.Type, entity.EntityName, method.Parameters)
		}

		if err := validateMethodArguments(method); err != nil {
			return err
		}

		if method.Type == enums.MethodTypeGetCollectionExclusive {
			if err := validateMethodExclusiveArguments(method); err != nil {
				return err
			}
		}

		spl := strings.Split(method.ConstructedParameters, ", ")
		method.ConstructedParametersFormatted = strings.Join(spl, ",\n                    ")

		if method.Type == enums.MethodTypeGetCollection ||
			method.Type == enums.MethodTypeGetCollectionPaged ||
			method.Type == enums.MethodTypeGetCount ||
			method.Type == enums.MethodTypeGetCollectionExclusive {
			handleCachePolicy(method, entity)
		}
	}

	return nil
}

func getTransformedExclusiveStartParameters(exclusiveStartParameters []*models.Parameter) []*models.Parameter {
	var transformedParameters = make([]*models.Parameter, 0)

	for _, parameter := range exclusiveStartParameters {
		transformedParameters = append(transformedParameters, &models.Parameter{
			Name: "ExclusiveStart" + toPascalCase(parameter.Name),
		})
	}

	return transformedParameters
}

func handleCachePolicy(method *models.Method, entity *models.Entity) {
	method.CachePolicy = "CacheManager.UnqualifiedNonExpiringCachePolicy"
	method.CollectionIdentifier = fmt.Sprintf("\"%s\"", method.Name)

	if len(method.Parameters) != 0 {
		lastIndex, key := determineKeyRaw(0, method.Parameters)
		_, cachePolicy := determineKey(0, method.Parameters, false)

		method.CollectionIdentifier = method.Name + "_" + key
		method.CachePolicy = "new CacheManager.CachePolicy(\n                CacheManager.CacheScopeFilter.Qualified,\n                " + cachePolicy + "\n            )"

		if v, ok := entity.StateTokenMap[key]; !ok || !v {
			entity.StateTokenMap[key] = true

			_, stateToken := determineKey(0, method.Parameters, true)
			entity.StateTokens = append(entity.StateTokens, fmt.Sprintf("new StateToken(%s)", stateToken))
		}

		if method.Type == enums.MethodTypeGetCollectionPaged {
			method.CollectionIdentifier += fmt.Sprintf("_StartRowIndex:{%d}_MaximumRows:{%d}", len(method.Parameters), len(method.Parameters)+1)
			method.CollectionIdentifier = fmt.Sprintf("string.Format(\"%s\", %sstartRowIndex, maximumRows)", method.CollectionIdentifier, method.ConstructedParameters)
		} else if method.Type == enums.MethodTypeGetCollectionExclusive {
			method.CollectionIdentifier += fmt.Sprintf("_Count:{%d}", len(method.Parameters))
			if len(method.ExclusiveStartParameters) > 0 {
				transformedParameters := getTransformedExclusiveStartParameters(method.ExclusiveStartParameters)
				_, exclusiveStartKey := determineKeyRaw(lastIndex+1, transformedParameters)
				method.CollectionIdentifier += "_" + exclusiveStartKey
				transformedParameters = nil
			}
			method.CollectionIdentifier += fmt.Sprintf("_ExclusiveStart%s:{%d}", entity.IDProperty.Name, len(method.Parameters)+len(method.ExclusiveStartParameters)+1)
			method.CollectionIdentifier = fmt.Sprintf("string.Format(\"%s\", %scount,%s exclusiveStart%s)", method.CollectionIdentifier, method.ConstructedParameters, method.ConstructedExclusiveStartParameters, helpers.NormalizePascalParts(entity.IDProperty.Name))
		} else {
			method.CollectionIdentifier = fmt.Sprintf("string.Format(\"%s\", %s)", method.CollectionIdentifier, method.ConstructedParameters)
		}
	}
}

func validateMethodExclusiveArguments(method *models.Method) error {
	for _, parameter := range method.ExclusiveStartParameters {
		if parameter.Name == "" {
			return ErrParameterNameNotSpecified
		}

		if parameter.Type == enums.CSharpTypeUnknown {
			parameter.Type = enums.CSharpTypeString
		}

		parameter.SqlDbType = enums.SqlDbTypeFromCSharpType(parameter.Type)

		if parameter.Type == enums.CSharpTypeDateTime {
			if parameter.IsUTC == nil {
				parameter.IsUTC = new(bool)
				*parameter.IsUTC = false
			}
		}

		if parameter.Type == enums.CSharpTypeString {
			if parameter.IsUnicode == nil {
				parameter.IsUnicode = new(bool)
				*parameter.IsUnicode = true
			}

			if *parameter.IsUnicode {
				parameter.SqlDbType = enums.SqlDbTypeNVarChar
			} else {
				parameter.SqlDbType = enums.SqlDbTypeVarChar
			}
		}

		if parameter.Type == enums.CSharpTypeByteArray {
			if parameter.IsVarBinary == nil {
				parameter.IsVarBinary = new(bool)
				*parameter.IsVarBinary = false
			}

			if !*parameter.IsVarBinary && parameter.Length == 0 {
				return ErrParameterNotVarBinaryWithNoLength
			}
		}

		method.ConstructedExclusiveStartStringParameters += parameter.Type.Name() + "? exclusiveStart" + toPascalCase(parameter.Name) + ", "
		method.ConstructedExclusiveStartParameters += "exclusiveStart" + toPascalCase(parameter.Name) + ", "
	}

	if len(method.ExclusiveStartParameters) > 0 {
		method.ConstructedExclusiveStartStringParameters = " " + method.ConstructedExclusiveStartStringParameters[:len(method.ConstructedExclusiveStartStringParameters)-1]
		method.ConstructedExclusiveStartParameters = " " + method.ConstructedExclusiveStartParameters[:len(method.ConstructedExclusiveStartParameters)-1]
	}

	return nil
}

func validateMethodArguments(method *models.Method) error {
	for i, parameter := range method.Parameters {
		if parameter.Name == "" {
			return ErrParameterNameNotSpecified
		}

		if parameter.Type == enums.CSharpTypeUnknown {
			parameter.Type = enums.CSharpTypeString
		}

		parameter.SqlDbType = enums.SqlDbTypeFromCSharpType(parameter.Type)

		if parameter.Type == enums.CSharpTypeDateTime {
			if parameter.IsUTC == nil {
				parameter.IsUTC = new(bool)
				*parameter.IsUTC = false
			}
		}

		if parameter.Type == enums.CSharpTypeString {
			if parameter.IsUnicode == nil {
				parameter.IsUnicode = new(bool)
				*parameter.IsUnicode = true
			}

			if *parameter.IsUnicode {
				parameter.SqlDbType = enums.SqlDbTypeNVarChar
			} else {
				parameter.SqlDbType = enums.SqlDbTypeVarChar
			}
		}

		if parameter.Type == enums.CSharpTypeByteArray {
			if parameter.IsVarBinary == nil {
				parameter.IsVarBinary = new(bool)
				*parameter.IsVarBinary = false
			}

			if !*parameter.IsVarBinary && parameter.Length == 0 {
				return ErrParameterNotVarBinaryWithNoLength
			}
		}

		if i+1 == len(method.Parameters) && method.Type != enums.MethodTypeGetCollectionPaged && method.Type != enums.MethodTypeGetCollectionExclusive {
			method.ConstructedStringParameters += parameter.Type.Name() + " " + parameter.Name
			method.ConstructedParameters += parameter.Name
		} else {
			method.ConstructedStringParameters += parameter.Type.Name() + " " + parameter.Name + ", "
			method.ConstructedParameters += parameter.Name + ", "
		}
	}

	return nil
}

func handleGetOrCreate(method *models.Method, entity *models.Entity) {
	_, method.LookupKey = determineKey(0, method.Parameters, false)
	_, rawKey := determineKeyRaw(0, method.Parameters)
	_, pascalizedKey := determineKey(0, method.Parameters, true)

	if v, ok := entity.LookupKeyMap[rawKey]; !ok || !v {
		entity.LookupKeyMap[rawKey] = true
		entity.LookupKeys = append(entity.LookupKeys, pascalizedKey)
	}
}

func handleLookup(method *models.Method, entity *models.Entity) error {
	if len(method.Parameters) == 0 {
		return ErrLookupMethodNoParameters
	}

	_, method.LookupKey = determineKey(0, method.Parameters, false)
	_, rawKey := determineKeyRaw(0, method.Parameters)
	_, pascalizedKey := determineKey(0, method.Parameters, true)

	if v, ok := entity.LookupKeyMap[rawKey]; !ok || !v {
		entity.LookupKeyMap[rawKey] = true
		entity.LookupKeys = append(entity.LookupKeys, pascalizedKey)
	}

	return nil
}

func validatePassiveProperties(method *models.Method, entity *models.Entity) error {
	for _, passiveProperty := range method.PassiveProperties {
		if passiveProperty.Name == "" {
			return ErrPassivePropertyNameNotSpecified
		}

		if passiveProperty.Value == "" {
			return ErrPassivePropertyValueNotSpecified
		}

		if passiveProperty.Name == entity.IDProperty.Name {
			return ErrPassivePropertyCannotBeIDProperty
		}

		if existsInArguments(method.Parameters, passiveProperty.Name) {
			return ErrPassivePropertyCannotBeMethodArgument
		}

		if exists, property := existsInProperties(entity.Properties, passiveProperty.Name); exists {
			passiveProperty.Value = parsePassiveValueBasedOnType(passiveProperty.Value, property.SqlDbType)

			continue
		}

		return ErrPassivePropertyIsNotProperty
	}

	return nil
}

func applySharedMethodData(method *models.Method, entity *models.Entity) {
	method.EntityName = entity.EntityName
	method.Table = entity.Table
	method.ConstructedStringParameters = ""
	method.IDProperty = entity.IDProperty
	method.RemoteCacheable = entity.RemoteCacheable
	method.Version = *entity.Version
	method.Properties = entity.Properties
	method.EntityType = entity.EntityType
}

func handleGetCount(method *models.Method) error {
	if method.CountReturnType == enums.CSharpTypeUnknown {
		method.CountReturnType = enums.CSharpTypeInt
	}

	if method.CountReturnType != enums.CSharpTypeInt &&
		method.CountReturnType != enums.CSharpTypeLong {
		return errors.New("count return type must be int or long")
	}

	return nil
}

func prevalidateMethod(method *models.Method) error {
	if method.Name == "" {
		return ErrMethodNameNotSpecified
	}

	if method.DALName == "" {
		method.DALName = method.Name
	}

	if method.Visibility == enums.VisibilityTypeUnknown {
		method.Visibility = enums.VisibilityTypePublic
	}

	if method.Parameters == nil {
		method.Parameters = make([]*models.Parameter, 0)
	}

	return nil
}

func validateProperties(entity *models.Entity) error {
	entity.EntityType = "Entity"
	hasCreated := false

	for _, property := range entity.Properties {
		if property.Name == "" {
			return ErrPropertyNameNotSpecified
		}

		if property.Type == enums.CSharpTypeUnknown {
			property.Type = enums.CSharpTypeString
		}

		if property.Visibility == enums.VisibilityTypeUnknown {
			property.Visibility = enums.VisibilityTypePublic
		}

		if property.SqlDbType == enums.SqlDbTypeUnknown {
			property.SqlDbType = enums.SqlDbTypeFromCSharpType(property.Type)
		}

		if property.Type == enums.CSharpTypeString {
			handleStringProperty(property)
		}

		if property.Type == enums.CSharpTypeDateTime {
			if property.IsUTC == nil {
				property.IsUTC = new(bool)
				*property.IsUTC = false
			}
		}

		if property.Type == enums.CSharpTypeByteArray {
			if err := handleByteArrayProperty(property); err != nil {
				return err
			}
		}

		if property.ReadOnly == nil {
			property.ReadOnly = new(bool)
			*property.ReadOnly = true
		}

		if property.Nullable == nil {
			property.Nullable = new(bool)
			*property.Nullable = false
		}

		if property.ForeignKey != "" && property.Name != "Created" && property.Name != "Updated" {
			if err := handleForeignKey(property, entity); err != nil {
				return err
			}
		}

		determineEntityType(property, entity, &hasCreated)
		parsePropertyArgsString(property, entity)
	}

	entity.PropertiesArgs = entity.PropertiesArgs[:len(entity.PropertiesArgs)-2]

	if !hasCreated {
		return ErrCreatedNotFoundOrWrongType
	}

	return nil
}

func parsePropertyArgsString(property *models.Property, entity *models.Entity) {
	if !*property.ReadOnly && property.Name != "Created" && property.Name != "Updated" {
		entity.PropertiesArgs += property.Type.Name()

		if *property.Nullable && property.Type != enums.CSharpTypeString {
			entity.PropertiesArgs += "?"
		}

		entity.PropertiesArgs += " " + toCamelCase(property.Name) + ", "
	}
}

func determineEntityType(property *models.Property, entity *models.Entity, hasCreated *bool) {
	if property.Name == "Updated" && property.Type == enums.CSharpTypeDateTime {
		entity.EntityType = "UpdateableEntity"

		entity.IsUpdatedNullable = *property.Nullable
	}

	if property.Name == "Created" && property.Type == enums.CSharpTypeDateTime {
		*hasCreated = true
	}
}

func handleForeignKey(property *models.Property, entity *models.Entity) error {
	if !ValidForeignKeySyntax.MatchString(property.ForeignKey) {
		return errors.New("invalid foreign key syntax")
	}

	parts := ValidForeignKeySyntax.FindStringSubmatch(property.ForeignKey)
	if parts[2] == "" {
		property.ConstructedForeignKey = fmt.Sprintf("[dbo].[%s] ([%s])", parts[3], parts[4])
	} else {
		property.ConstructedForeignKey = fmt.Sprintf("[%s].[dbo].[%s] ([%s])", parts[2], parts[3], parts[4])
	}

	property.ConstructedForeignKeyConstraintKey = fmt.Sprintf("FK_%s_%s_%s", entity.Table, parts[3], property.Name)

	return nil
}

func handleByteArrayProperty(property *models.Property) error {
	if property.IsVarBinary == nil {
		property.IsVarBinary = new(bool)
		*property.IsVarBinary = false
	}

	if !*property.IsVarBinary && property.Length == 0 {
		return ErrPropertyNotVarBinaryWithNoLength
	}

	if *property.IsVarBinary {
		property.SqlDbType = enums.SqlDbTypeVarBinary
	} else {
		property.SqlDbType = enums.SqlDbTypeBinary
	}

	return nil
}

func handleStringProperty(property *models.Property) {
	if property.IsUnicode == nil {
		property.IsUnicode = new(bool)
		*property.IsUnicode = true
	}

	if *property.IsUnicode {
		property.SqlDbType = enums.SqlDbTypeNVarChar
	} else {
		property.SqlDbType = enums.SqlDbTypeVarChar
	}
}

func validatePredefined(entity *models.Entity) error {
	if entity.Predefined != nil {
		if entity.Predefined.GetMethod == "" {
			return ErrPredefinedGetMethodName
		}

		for _, value := range entity.Predefined.Values {
			if value.Name == "" {
				return ErrPredefinedValueName
			}

			if value.Properties == nil || len(value.Properties) == 0 {
				value.Properties = make([]map[string]string, 1)
				value.Properties[0] = make(map[string]string)

				if strings.HasPrefix(entity.Predefined.GetMethod, "GetBy") {

					argumentName := strings.TrimPrefix(entity.Predefined.GetMethod, "GetBy")
					argumentNameSplit := strings.Split(argumentName, "And")

					if len(argumentNameSplit) > 1 {
						return ErrCannotDeterminePredefinedLookupKeyWithMultipleArguments
					}

					value.Properties[0][argumentNameSplit[0]] = value.Name
				} else {
					return ErrCannotDeterminePredefinedLookupKey
				}
			}

			value.PropertiesConstructed = fmt.Sprintf("%s%s = %s(", value.Name, entity.IDProperty.Name, entity.Predefined.GetMethod)

			for _, v := range value.Properties {

				key, _ := getFirstElement(v)
				fieldName := value.Name + toPascalCase(key)

				value.PropertiesConstructed += fieldName + ", "
			}

			value.PropertiesConstructed = value.PropertiesConstructed[:len(value.PropertiesConstructed)-2] + ")." + entity.IDProperty.Name
		}
	}

	return nil
}

func setInitialLookups(entity *models.Entity) {
	entity.LookupKeys = make([]string, 0)
	entity.StateTokens = make([]string, 0)
	entity.LookupKeyMap = make(map[string]bool)
	entity.StateTokenMap = make(map[string]bool)
}

func validateRemoteCacheable(entity *models.Entity) error {
	if entity.RemoteCacheable != nil && entity.RemoteCacheable.MemcachedGroupSetting != "" {

		parts := strings.Split(entity.RemoteCacheable.MemcachedGroupSetting, ":")
		if len(parts) > 2 {
			return ErrInvalidMemcachedGroupName
		}

		entity.RemoteCacheable.MemcachedGroupSettingClass = parts[0]

		if len(parts) == 2 {
			entity.RemoteCacheable.MemcachedGroupSettingProperty = parts[1]
		}
	}

	return nil
}

func validateIdProperty(entity *models.Entity) error {
	if entity.IDProperty.Name == "" {
		entity.IDProperty.Name = "ID"
	}

	if entity.IDProperty.Type == enums.CSharpTypeUnknown {
		entity.IDProperty.Type = enums.CSharpTypeInt
	}

	if entity.IDProperty.Visibility == enums.VisibilityTypeUnknown {
		entity.IDProperty.Visibility = enums.VisibilityTypePublic
	}

	if entity.IDProperty.ReadOnly == nil {
		entity.IDProperty.ReadOnly = new(bool)
		*entity.IDProperty.ReadOnly = true
	}

	if entity.IDProperty.Nullable == nil {
		entity.IDProperty.Nullable = new(bool)
		*entity.IDProperty.Nullable = false
	}

	if entity.IDProperty.Type != enums.CSharpTypeByte &&
		entity.IDProperty.Type != enums.CSharpTypeInt &&
		entity.IDProperty.Type != enums.CSharpTypeLong {
		return ErrIDPropertyNotNumberType
	}

	entity.IDProperty.SqlDbType = enums.SqlDbTypeFromCSharpType(entity.IDProperty.Type)

	return nil
}

func preValidate(entity *models.Entity) error {
	if entity == nil {
		return ErrNilEntity
	}

	if entity.EntityName == "" {
		return ErrEntityNameNotSpecified
	}

	if entity.Database == "" {
		return ErrDatabaseNotSpecified
	}

	if entity.Table == "" {
		return ErrTableNotSpecified
	}

	if entity.IDProperty == nil {
		return ErrIDPropertyNotSpecified
	}

	if len(entity.Properties) == 0 {
		return ErrPropertiesNotSpecified
	}

	if entity.Visibility == enums.VisibilityTypeUnknown {
		entity.Visibility = enums.VisibilityTypeInternal
	}

	if entity.DALVisibility == enums.VisibilityTypeUnknown {
		entity.DALVisibility = enums.VisibilityTypeInternal
	}

	if entity.Version == nil {
		entity.Version = new(int)
		*entity.Version = 2
	}

	if *entity.Version != 1 && *entity.Version != 2 {
		return ErrBadVersionNumber
	}

	if entity.Namespace == "" {
		entity.Namespace = "Roblox"
	}

	if entity.DALNamespace == "" {
		entity.DALNamespace = entity.Namespace
	}

	if entity.GenerateMustGet == nil {
		entity.GenerateMustGet = new(bool)
		*entity.GenerateMustGet = false
	}

	if entity.GenerateCreateNew == nil {
		entity.GenerateCreateNew = new(bool)
		*entity.GenerateCreateNew = false
	}

	return nil
}

func containsDataTables(dataTables []enums.SqlDbType, sqlDbType enums.SqlDbType) bool {
	for _, dt := range dataTables {
		if dt == sqlDbType {
			return true
		}
	}

	return false
}

func buildDatabasesBasedOnEntities(entities []*models.Entity, currentDatabases map[string]*models.Database) map[string]*models.Database {
	databases := make(map[string]*models.Database)

	for _, entity := range entities {
		if _, ok := databases[entity.Database]; !ok {
			if db, ok := currentDatabases[entity.Database]; ok {
				databases[entity.Database] = db
			} else {
				databases[entity.Database] = &models.Database{
					Name:       entity.Database,
					DataTables: make([]enums.SqlDbType, 0),
					Entities:   make([]*models.Entity, 0),
				}
			}
		}

		if databases[entity.Database].DatabaseSharding != nil {
			sharding := databases[entity.Database].DatabaseSharding

			if sharding.IsWindows == nil {
				sharding.IsWindows = new(bool)
				*sharding.IsWindows = false
			}

			if sharding.ShardCount == nil {
				sharding.ShardCount = new(int)
				*sharding.ShardCount = 8
			}

			if sharding.InitialSize == "" {
				sharding.InitialSize = "5120KB"
			}

			if sharding.MaximumSize == "" {
				sharding.MaximumSize = "10MB"
			}

			if sharding.GrowthSize == "" {
				sharding.GrowthSize = "1024KB"
			}
		}

		for _, method := range entity.Methods {
			if method.Type == enums.MethodTypeMultiGet {
				sqlDbType := enums.SqlDbTypeFromCSharpType(entity.IDProperty.Type)

				if !containsDataTables(databases[entity.Database].DataTables, sqlDbType) {
					databases[entity.Database].DataTables = append(databases[entity.Database].DataTables, sqlDbType)
				}
			}
		}
	}

	// We need to order the entities by the order where foreign keys are defined.
	// i.e. if agents has a foreign key to agent types, then agent types must be defined first.
	// This is because the foreign key must be defined after the table it references.

	// We will build a dependency graph to determine the order of the entities.
	// We will use a map to store the dependencies.

	var dependencies = make(map[string][]string)

	for _, entity := range entities {
		dependencies[entity.Table] = make([]string, 0)
	}

	for _, entity := range entities {
		for _, property := range entity.Properties {
			if property.ForeignKey != "" {
				parts := ValidForeignKeySyntax.FindStringSubmatch(property.ForeignKey)
				dependencies[entity.Table] = append(dependencies[entity.Table], parts[3])
			}
		}
	}

	// Now we will sort the entities based on the dependencies.
	// i.e. EntityA -> EntityB, EntityC; therefore EntityB and EntityC must be defined first.

	var sortedEntities = make([]*models.Entity, 0)
	var visited = make(map[string]bool)

	var visitAll func(string)
	visitAll = func(entityName string) {
		if visited[entityName] {
			return
		}

		visited[entityName] = true

		for _, dependentEntity := range dependencies[entityName] {
			visitAll(dependentEntity)
		}

		for _, entity := range entities {
			if entity.Table == entityName {
				sortedEntities = append(sortedEntities, entity)
			}
		}
	}

	for _, entity := range entities {
		visitAll(entity.Table)
	}

	// Now replace the entities with the sorted entities on each database.
	for _, database := range databases {
		for _, entity := range sortedEntities {
			if entity.Database == database.Name {
				database.Entities = append(database.Entities, entity)
			}
		}
	}

	for _, database := range currentDatabases {
		if _, ok := databases[database.Name]; !ok {
			databases[database.Name] = database // case here for empty databases.
		}
	}

	return databases
}

func determineIfDatabaseConfigFile(fileName string) bool {
	// Should be 3 parts:
	// {Name}.database.{ext}
	// e.g. MyDatabase.database.json or MyDatabase.database.yaml
	// if only 2 parts, or the middle part is not "database", then it is an entity file.
	parts := strings.Split(fileName, ".")
	if len(parts) != 3 {
		return false
	}

	if parts[1] != "database" {
		return false
	}

	return true
}

func validateDatabase(database *models.Database) error {
	if database.Name == "" {
		return ErrDatabaseNotSpecified
	}

	if database.DatabaseSharding != nil {
		sharding := database.DatabaseSharding

		if sharding.IsWindows == nil {
			sharding.IsWindows = new(bool)
			*sharding.IsWindows = false
		}

		if sharding.ShardCount == nil {
			sharding.ShardCount = new(int)
			*sharding.ShardCount = 8
		}

		if sharding.InitialSize == "" {
			sharding.InitialSize = "5120KB"
		}

		if sharding.MaximumSize == "" {
			sharding.MaximumSize = "10MB"
		}

		if sharding.GrowthSize == "" {
			sharding.GrowthSize = "1024KB"
		}
	}

	return nil
}

// Parse parses the configuration files.
// And produces the full configuration model.
func Parse() ([]*models.Entity, map[string]*models.Database, error) {
	if *flags.ConfigurationDirectoryFlag == "" {
		return nil, nil, ErrConfigurationDirectoryNotSpecified
	}

	var entities []*models.Entity
	var databases map[string]*models.Database = make(map[string]*models.Database)

	err := filepath.Walk(*flags.ConfigurationDirectoryFlag, func(fileName string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fileInfo.IsDir() && !*flags.RecurseFlag {
			return filepath.SkipDir
		}

		if !fileInfo.IsDir() && !strings.HasSuffix(fileName, ".json") && !strings.HasSuffix(fileName, ".yaml") {
			if !determineIfDatabaseConfigFile(fileInfo.Name()) {
				println("Parsing entity file: ", fileName)
				entity, err := parseEntityFileDependingOnExtension(fileName)
				if err == io.EOF {
					// Skip the file.
					return nil
				}
				if err != nil {
					return err
				}

				if entity != nil {
					entities = append(entities, entity)
				}
			} else {
				println("Parsing database file: ", fileName)
				database, err := parseDatabaseFileDependingOnExtension(fileName)
				if err != nil {
					return err
				}

				if database != nil {
					if err := validateDatabase(database); err != nil {
						return err
					}

					databases[database.Name] = database
				}
			}
		}

		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	for _, entity := range entities {
		println("Validating entity: ", entity.EntityName)
		if err := validateNewEntity(entity); err != nil {
			return nil, nil, err
		}
	}

	println("Building databases based on entities")
	databases = buildDatabasesBasedOnEntities(entities, databases)

	return entities, databases, nil
}
