// Code generated by "stringer -type=Role"; DO NOT EDIT.

package uast

import "fmt"

const _Role_name = "SimpleIdentifierQualifiedIdentifierBinaryExpressionBinaryExpressionLeftBinaryExpressionRightBinaryExpressionOpOpBitwiseLeftShiftOpBitwiseRightShiftOpBitwiseUnsignedRightShiftOpBitwiseOrOpBitwiseXorOpBitwiseAndExpressionStatementOpEqualOpNotEqualOpLessThanOpLessThanEqualOpGreaterThanOpGreaterThanEqualOpSameOpNotSameOpContainsOpNotContainsOpPreIncrementOpPostIncrementOpPreDecrementOpPostDecrementOpNegativeOpPositiveOpBitwiseComplementOpDereferenceOpTakeAddressFileOpBooleanAndOpBooleanOrOpBooleanNotOpBooleanXorOpAddOpSubstractOpMultiplyOpDivideOpModPackageDeclarationImportDeclarationImportPathImportAliasFunctionDeclarationFunctionDeclarationBodyFunctionDeclarationNameFunctionDeclarationReceiverFunctionDeclarationArgumentFunctionDeclarationArgumentNameFunctionDeclarationArgumentDefaultValueFunctionDeclarationVarArgsListTypeDeclarationTypeDeclarationBodyTypeDeclarationBasesTypeDeclarationImplementsVisibleFromInstanceVisibleFromTypeVisibleFromSubtypeVisibleFromPackageVisibleFromSubpackageVisibleFromModuleVisibleFromFriendVisibleFromWorldIfIfConditionIfBodyIfElseSwitchSwitchCaseSwitchCaseConditionSwitchCaseBodySwitchDefaultForForInitForExpressionForUpdateForBodyForEachWhileWhileConditionWhileBodyDoWhileDoWhileConditionDoWhileBodyBreakContinueBlockBlockScopeReturnTryTryBodyTryCatchTryFinallyThrowAssertCallCallReceiverCallCalleeCallPositionalArgumentCallNamedArgumentCallNamedArgumentNameCallNamedArgumentValueNoopBooleanLiteralByteLiteralByteStringLiteralCharacterLiteralListLiteralMapLiteralNullLiteralNumberLiteralRegexpLiteralSetLiteralStringLiteralTupleLiteralTypeLiteralOtherLiteralMapEntryMapKeyMapValueTypePrimitiveTypeAssignmentAssignmentVariableAssignmentValueAugmentedAssignmentAugmentedAssignmentOperatorAugmentedAssignmentVariableAugmentedAssignmentValueThisCommentDocumentationWhitespace"

var _Role_index = [...]uint16{0, 16, 35, 51, 71, 92, 110, 128, 147, 174, 185, 197, 209, 219, 228, 235, 245, 255, 270, 283, 301, 307, 316, 326, 339, 353, 368, 382, 397, 407, 417, 436, 449, 462, 466, 478, 489, 501, 513, 518, 529, 539, 547, 552, 570, 587, 597, 608, 627, 650, 673, 700, 727, 758, 797, 827, 842, 861, 881, 906, 925, 940, 958, 976, 997, 1014, 1031, 1047, 1049, 1060, 1066, 1072, 1078, 1088, 1107, 1121, 1134, 1137, 1144, 1157, 1166, 1173, 1180, 1185, 1199, 1208, 1215, 1231, 1242, 1247, 1255, 1260, 1270, 1276, 1279, 1286, 1294, 1304, 1309, 1315, 1319, 1331, 1341, 1363, 1380, 1401, 1423, 1427, 1441, 1452, 1469, 1485, 1496, 1506, 1517, 1530, 1543, 1553, 1566, 1578, 1589, 1601, 1609, 1615, 1623, 1627, 1640, 1650, 1668, 1683, 1702, 1729, 1756, 1780, 1784, 1791, 1804, 1814}

func (i Role) String() string {
	i -= 1
	if i < 0 || i >= Role(len(_Role_index)-1) {
		return fmt.Sprintf("Role(%d)", i+1)
	}
	return _Role_name[_Role_index[i]:_Role_index[i+1]]
}
