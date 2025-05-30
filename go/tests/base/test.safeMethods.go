package base
import "github.com/ccxt/ccxt/go/v4"

// PLEASE DO NOT EDIT THIS FILE, IT IS GENERATED AND WILL BE OVERWRITTEN:
// https://github.com/ccxt/ccxt/blob/master/CONTRIBUTING.md#how-to-contribute-code

func TestSafeMethods()  {
    exchange := ccxt.NewExchange().(*ccxt.Exchange); exchange.DerivedExchange = exchange; exchange.InitParent(map[string]interface{} {
        "id": "regirock",
    }, map[string]interface{}{}, exchange)
    var inputDict map[string]interface{} = map[string]interface{} {
        "i": 1,
        "f": 0.123,
        "bool": true,
        "list": []interface{}{1, 2, 3},
        "dict": map[string]interface{} {
            "a": 1,
        },
        "str": "heLlo",
        "strNumber": "3",
        "zeroNumeric": 0,
        "zeroString": "0",
        "undefined": nil,
        "emptyString": "",
        "floatNumeric": 0.123,
        "floatString": "0.123",
    }
    var inputList interface{} = []interface{}{"Hi", 2}
    var compareDict map[string]interface{} = map[string]interface{} {
        "a": 1,
    }
    var compareList interface{} = []interface{}{1, 2, 3}
    var factor interface{} = 10
    // safeValue
    Assert(IsEqual(exchange.SafeValue(inputDict, "i"), 1))
    Assert(IsEqual(exchange.SafeValue(inputDict, "f"), 0.123))
    Assert(IsEqual(exchange.SafeValue(inputDict, "bool"), true))
    Assert(Equals(exchange.SafeValue(inputDict, "list"), compareList))
    var dictObject interface{} = exchange.SafeValue(inputDict, "dict")
    Assert(Equals(dictObject, compareDict))
    Assert(IsEqual(exchange.SafeValue(inputDict, "str"), "heLlo"))
    Assert(IsEqual(exchange.SafeValue(inputDict, "strNumber"), "3"))
    Assert(IsEqual(exchange.SafeValue(inputList, 0), "Hi"))
    // safeValue2
    Assert(IsEqual(exchange.SafeValue2(inputDict, "a", "i"), 1))
    Assert(IsEqual(exchange.SafeValue2(inputDict, "a", "f"), 0.123))
    Assert(IsEqual(exchange.SafeValue2(inputDict, "a", "bool"), true))
    Assert(Equals(exchange.SafeValue2(inputDict, "a", "list"), compareList))
    dictObject = exchange.SafeValue2(inputDict, "a", "dict")
    Assert(Equals(dictObject, compareDict))
    Assert(IsEqual(exchange.SafeValue2(inputDict, "a", "str"), "heLlo"))
    Assert(IsEqual(exchange.SafeValue2(inputDict, "a", "strNumber"), "3"))
    Assert(IsEqual(exchange.SafeValue2(inputList, 2, 0), "Hi"))
    // safeValueN
    Assert(IsEqual(exchange.SafeValueN(inputDict, []interface{}{"a", "b", "i"}), 1))
    Assert(IsEqual(exchange.SafeValueN(inputDict, []interface{}{"a", "b", "f"}), 0.123))
    Assert(IsEqual(exchange.SafeValueN(inputDict, []interface{}{"a", "b", "bool"}), true))
    Assert(Equals(exchange.SafeValueN(inputDict, []interface{}{"a", "b", "list"}), compareList))
    dictObject = exchange.SafeValueN(inputDict, []interface{}{"a", "b", "dict"})
    Assert(Equals(dictObject, compareDict))
    Assert(IsEqual(exchange.SafeValueN(inputDict, []interface{}{"a", "b", "str"}), "heLlo"))
    Assert(IsEqual(exchange.SafeValueN(inputDict, []interface{}{"a", "b", "strNumber"}), "3"))
    Assert(IsEqual(exchange.SafeValueN(inputList, []interface{}{3, 2, 0}), "Hi"))
    // safeDict
    dictObject = exchange.SafeDict(inputDict, "dict")
    Assert(Equals(dictObject, compareDict))
    var listObject interface{} = exchange.SafeDict(inputDict, "list")
    Assert(IsEqual(listObject, nil))
    Assert(IsEqual(exchange.SafeDict(inputList, 1), nil))
    // safeDict2
    dictObject = exchange.SafeDict2(inputDict, "a", "dict")
    Assert(Equals(dictObject, compareDict))
    listObject = exchange.SafeDict2(inputDict, "a", "list")
    Assert(IsEqual(listObject, nil))
    // @ts-expect-error
    Assert(IsEqual(exchange.SafeDict2(inputList, 2, 1), nil))
    // safeDictN
    dictObject = exchange.SafeDictN(inputDict, []interface{}{"a", "b", "dict"})
    Assert(Equals(dictObject, compareDict))
    listObject = exchange.SafeDictN(inputDict, []interface{}{"a", "b", "list"})
    Assert(IsEqual(listObject, nil))
    Assert(IsEqual(exchange.SafeDictN(inputList, []interface{}{3, 2, 1}), nil))
    // safeList
    listObject = exchange.SafeList(inputDict, "list")
    Assert(Equals(dictObject, compareDict))
    Assert(IsEqual(exchange.SafeList(inputDict, "dict"), nil))
    Assert(IsEqual(exchange.SafeList(inputList, 1), nil))
    // safeList2
    listObject = exchange.SafeList2(inputDict, "a", "list")
    Assert(Equals(dictObject, compareDict))
    Assert(IsEqual(exchange.SafeList2(inputDict, "a", "dict"), nil))
    // @ts-expect-error
    Assert(IsEqual(exchange.SafeList2(inputList, 2, 1), nil))
    // safeListN
    listObject = exchange.SafeListN(inputDict, []interface{}{"a", "b", "list"})
    Assert(Equals(dictObject, compareDict))
    Assert(IsEqual(exchange.SafeListN(inputDict, []interface{}{"a", "b", "dict"}), nil))
    Assert(IsEqual(exchange.SafeListN(inputList, []interface{}{3, 2, 1}), nil))
    // safeString
    Assert(IsEqual(exchange.SafeString(inputDict, "i"), "1"))
    Assert(IsEqual(exchange.SafeString(inputDict, "f"), "0.123"))
    // assert (exchange.safeString (inputDict, 'bool') === 'true'); returns True in python and 'true' in js
    Assert(IsEqual(exchange.SafeString(inputDict, "str"), "heLlo"))
    Assert(IsEqual(exchange.SafeString(inputDict, "strNumber"), "3"))
    Assert(IsEqual(exchange.SafeString(inputList, 0), "Hi"))
    // safeString2
    Assert(IsEqual(exchange.SafeString2(inputDict, "a", "i"), "1"))
    Assert(IsEqual(exchange.SafeString2(inputDict, "a", "f"), "0.123"))
    Assert(IsEqual(exchange.SafeString2(inputDict, "a", "str"), "heLlo"))
    Assert(IsEqual(exchange.SafeString2(inputDict, "a", "strNumber"), "3"))
    Assert(IsEqual(exchange.SafeString2(inputList, 2, 0), "Hi"))
    // safeStringN
    Assert(IsEqual(exchange.SafeStringN(inputDict, []interface{}{"a", "b", "i"}), "1"))
    Assert(IsEqual(exchange.SafeStringN(inputDict, []interface{}{"a", "b", "f"}), "0.123"))
    Assert(IsEqual(exchange.SafeStringN(inputDict, []interface{}{"a", "b", "str"}), "heLlo"))
    Assert(IsEqual(exchange.SafeStringN(inputDict, []interface{}{"a", "b", "strNumber"}), "3"))
    Assert(IsEqual(exchange.SafeStringN(inputList, []interface{}{3, 2, 0}), "Hi"))
    // safeStringLower
    Assert(IsEqual(exchange.SafeStringLower(inputDict, "i"), "1"))
    Assert(IsEqual(exchange.SafeStringLower(inputDict, "f"), "0.123"))
    Assert(IsEqual(exchange.SafeStringLower(inputDict, "str"), "hello"))
    Assert(IsEqual(exchange.SafeStringLower(inputDict, "strNumber"), "3"))
    Assert(IsEqual(exchange.SafeStringLower(inputList, 0), "hi"))
    // safeStringLower2
    Assert(IsEqual(exchange.SafeStringLower2(inputDict, "a", "i"), "1"))
    Assert(IsEqual(exchange.SafeStringLower2(inputDict, "a", "f"), "0.123"))
    Assert(IsEqual(exchange.SafeStringLower2(inputDict, "a", "str"), "hello"))
    Assert(IsEqual(exchange.SafeStringLower2(inputDict, "a", "strNumber"), "3"))
    Assert(IsEqual(exchange.SafeStringLower2(inputList, 2, 0), "hi"))
    // safeStringLowerN
    Assert(IsEqual(exchange.SafeStringLowerN(inputDict, []interface{}{"a", "b", "i"}), "1"))
    Assert(IsEqual(exchange.SafeStringLowerN(inputDict, []interface{}{"a", "b", "f"}), "0.123"))
    Assert(IsEqual(exchange.SafeStringLowerN(inputDict, []interface{}{"a", "b", "str"}), "hello"))
    Assert(IsEqual(exchange.SafeStringLowerN(inputDict, []interface{}{"a", "b", "strNumber"}), "3"))
    Assert(IsEqual(exchange.SafeStringLowerN(inputList, []interface{}{3, 2, 0}), "hi"))
    // safeStringUpper
    Assert(IsEqual(exchange.SafeStringUpper(inputDict, "i"), "1"))
    Assert(IsEqual(exchange.SafeStringUpper(inputDict, "f"), "0.123"))
    Assert(IsEqual(exchange.SafeStringUpper(inputDict, "str"), "HELLO"))
    Assert(IsEqual(exchange.SafeStringUpper(inputDict, "strNumber"), "3"))
    Assert(IsEqual(exchange.SafeStringUpper(inputList, 0), "HI"))
    // safeStringUpper2
    Assert(IsEqual(exchange.SafeStringUpper2(inputDict, "a", "i"), "1"))
    Assert(IsEqual(exchange.SafeStringUpper2(inputDict, "a", "f"), "0.123"))
    Assert(IsEqual(exchange.SafeStringUpper2(inputDict, "a", "str"), "HELLO"))
    Assert(IsEqual(exchange.SafeStringUpper2(inputDict, "a", "strNumber"), "3"))
    Assert(IsEqual(exchange.SafeStringUpper2(inputList, 2, 0), "HI"))
    // safeStringUpperN
    Assert(IsEqual(exchange.SafeStringUpperN(inputDict, []interface{}{"a", "b", "i"}), "1"))
    Assert(IsEqual(exchange.SafeStringUpperN(inputDict, []interface{}{"a", "b", "f"}), "0.123"))
    Assert(IsEqual(exchange.SafeStringUpperN(inputDict, []interface{}{"a", "b", "str"}), "HELLO"))
    Assert(IsEqual(exchange.SafeStringUpperN(inputDict, []interface{}{"a", "b", "strNumber"}), "3"))
    Assert(IsEqual(exchange.SafeStringUpperN(inputList, []interface{}{3, 2, 0}), "HI"))
    // safeInteger
    Assert(IsEqual(exchange.SafeInteger(inputDict, "i"), 1))
    Assert(IsEqual(exchange.SafeInteger(inputDict, "f"), 0))
    Assert(IsEqual(exchange.SafeInteger(inputDict, "strNumber"), 3))
    Assert(IsEqual(exchange.SafeInteger(inputList, 1), 2))
    // safeInteger2
    Assert(IsEqual(exchange.SafeInteger2(inputDict, "a", "i"), 1))
    Assert(IsEqual(exchange.SafeInteger2(inputDict, "a", "f"), 0))
    Assert(IsEqual(exchange.SafeInteger2(inputDict, "a", "strNumber"), 3))
    Assert(IsEqual(exchange.SafeInteger2(inputList, 2, 1), 2))
    // safeIntegerN
    Assert(IsEqual(exchange.SafeIntegerN(inputDict, []interface{}{"a", "b", "i"}), 1))
    Assert(IsEqual(exchange.SafeIntegerN(inputDict, []interface{}{"a", "b", "f"}), 0))
    Assert(IsEqual(exchange.SafeIntegerN(inputDict, []interface{}{"a", "b", "strNumber"}), 3))
    Assert(IsEqual(exchange.SafeIntegerN(inputList, []interface{}{3, 2, 1}), 2))
    // safeIntegerOmitZero
    Assert(IsEqual(exchange.SafeIntegerOmitZero(inputDict, "i"), 1))
    Assert(IsEqual(exchange.SafeIntegerOmitZero(inputDict, "f"), nil))
    Assert(IsEqual(exchange.SafeIntegerOmitZero(inputDict, "strNumber"), 3))
    Assert(IsEqual(exchange.SafeIntegerOmitZero(inputList, 1), 2))
    // safeIntegerProduct
    Assert(IsEqual(exchange.SafeIntegerProduct(inputDict, "i", factor), 10))
    Assert(IsEqual(exchange.SafeIntegerProduct(inputDict, "f", factor), 1)) // NB the result is 1
    Assert(IsEqual(exchange.SafeIntegerProduct(inputDict, "strNumber", factor), 30))
    Assert(IsEqual(exchange.SafeIntegerProduct(inputList, 1, factor), 20))
    // safeIntegerProduct2
    Assert(IsEqual(exchange.SafeIntegerProduct2(inputDict, "a", "i", factor), 10))
    Assert(IsEqual(exchange.SafeIntegerProduct2(inputDict, "a", "f", factor), 1)) // NB the result is 1
    Assert(IsEqual(exchange.SafeIntegerProduct2(inputDict, "a", "strNumber", factor), 30))
    Assert(IsEqual(exchange.SafeIntegerProduct2(inputList, 2, 1, factor), 20))
    // safeIntegerProductN
    Assert(IsEqual(exchange.SafeIntegerProductN(inputDict, []interface{}{"a", "b", "i"}, factor), 10))
    Assert(IsEqual(exchange.SafeIntegerProductN(inputDict, []interface{}{"a", "b", "f"}, factor), 1)) // NB the result is 1
    Assert(IsEqual(exchange.SafeIntegerProductN(inputDict, []interface{}{"a", "b", "strNumber"}, factor), 30))
    Assert(IsEqual(exchange.SafeIntegerProductN(inputList, []interface{}{3, 2, 1}, factor), 20))
    // safeTimestamp
    Assert(IsEqual(exchange.SafeTimestamp(inputDict, "i"), 1000))
    Assert(IsEqual(exchange.SafeTimestamp(inputDict, "f"), 123))
    Assert(IsEqual(exchange.SafeTimestamp(inputDict, "strNumber"), 3000))
    Assert(IsEqual(exchange.SafeTimestamp(inputList, 1), 2000))
    // safeTimestamp2
    Assert(IsEqual(exchange.SafeTimestamp2(inputDict, "a", "i"), 1000))
    Assert(IsEqual(exchange.SafeTimestamp2(inputDict, "a", "f"), 123))
    Assert(IsEqual(exchange.SafeTimestamp2(inputDict, "a", "strNumber"), 3000))
    Assert(IsEqual(exchange.SafeTimestamp2(inputList, 2, 1), 2000))
    // safeTimestampN
    Assert(IsEqual(exchange.SafeTimestampN(inputDict, []interface{}{"a", "b", "i"}), 1000))
    Assert(IsEqual(exchange.SafeTimestampN(inputDict, []interface{}{"a", "b", "f"}), 123))
    Assert(IsEqual(exchange.SafeTimestampN(inputDict, []interface{}{"a", "b", "strNumber"}), 3000))
    Assert(IsEqual(exchange.SafeTimestampN(inputList, []interface{}{3, 2, 1}), 2000))
    // safeFloat
    // @ts-expect-error
    Assert(IsEqual(exchange.SafeFloat(inputDict, "i"), ParseFloat(1)))
    Assert(IsEqual(exchange.SafeFloat(inputDict, "f"), 0.123))
    // @ts-expect-error
    Assert(IsEqual(exchange.SafeFloat(inputDict, "strNumber"), ParseFloat(3)))
    // @ts-expect-error
    Assert(IsEqual(exchange.SafeFloat(inputList, 1), ParseFloat(2)))
    // safeFloat2
    // @ts-expect-error
    Assert(IsEqual(exchange.SafeFloat2(inputDict, "a", "i"), ParseFloat(1)))
    Assert(IsEqual(exchange.SafeFloat2(inputDict, "a", "f"), 0.123))
    // @ts-expect-error
    Assert(IsEqual(exchange.SafeFloat2(inputDict, "a", "strNumber"), ParseFloat(3)))
    // @ts-expect-error
    Assert(IsEqual(exchange.SafeFloat2(inputList, 2, 1), ParseFloat(2)))
    // safeFloatN
    // @ts-expect-error
    Assert(IsEqual(exchange.SafeFloatN(inputDict, []interface{}{"a", "b", "i"}), ParseFloat(1)))
    Assert(IsEqual(exchange.SafeFloatN(inputDict, []interface{}{"a", "b", "f"}), 0.123))
    // @ts-expect-error
    Assert(IsEqual(exchange.SafeFloatN(inputDict, []interface{}{"a", "b", "strNumber"}), ParseFloat(3)))
    // @ts-expect-error
    Assert(IsEqual(exchange.SafeFloatN(inputList, []interface{}{3, 2, 1}), ParseFloat(2)))
    // safeNumber
    Assert(IsEqual(exchange.SafeNumber(inputDict, "i"), exchange.ParseNumber(1)))
    Assert(IsEqual(exchange.SafeNumber(inputDict, "f"), exchange.ParseNumber(0.123)))
    Assert(IsEqual(exchange.SafeNumber(inputDict, "strNumber"), exchange.ParseNumber(3)))
    Assert(IsEqual(exchange.SafeNumber(inputList, 1), exchange.ParseNumber(2)))
    Assert(IsEqual(exchange.SafeNumber(inputList, "bool"), nil))
    Assert(IsEqual(exchange.SafeNumber(inputList, "list"), nil))
    Assert(IsEqual(exchange.SafeNumber(inputList, "dict"), nil))
    Assert(IsEqual(exchange.SafeNumber(inputList, "str"), nil))
    // safeNumber2
    Assert(IsEqual(exchange.SafeNumber2(inputDict, "a", "i"), exchange.ParseNumber(1)))
    Assert(IsEqual(exchange.SafeNumber2(inputDict, "a", "f"), exchange.ParseNumber(0.123)))
    Assert(IsEqual(exchange.SafeNumber2(inputDict, "a", "strNumber"), exchange.ParseNumber(3)))
    Assert(IsEqual(exchange.SafeNumber2(inputList, 2, 1), exchange.ParseNumber(2)))
    // safeNumberN
    Assert(IsEqual(exchange.SafeNumberN(inputDict, []interface{}{"a", "b", "i"}), exchange.ParseNumber(1)))
    Assert(IsEqual(exchange.SafeNumberN(inputDict, []interface{}{"a", "b", "f"}), exchange.ParseNumber(0.123)))
    Assert(IsEqual(exchange.SafeNumberN(inputDict, []interface{}{"a", "b", "strNumber"}), exchange.ParseNumber(3)))
    Assert(IsEqual(exchange.SafeNumberN(inputList, []interface{}{3, 2, 1}), exchange.ParseNumber(2)))
    // safeBool
    Assert(IsEqual(exchange.SafeBool(inputDict, "bool"), true))
    Assert(IsEqual(exchange.SafeBool(inputList, 1), nil))
    // safeBool2
    Assert(IsEqual(exchange.SafeBool2(inputDict, "a", "bool"), true))
    Assert(IsEqual(exchange.SafeBool2(inputList, 2, 1), nil))
    // safeBoolN
    Assert(IsEqual(exchange.SafeBoolN(inputDict, []interface{}{"a", "b", "bool"}), true))
    Assert(IsEqual(exchange.SafeBoolN(inputList, []interface{}{3, 2, 1}), nil))
    // safeNumberOmitZero
    Assert(IsEqual(exchange.SafeNumberOmitZero(inputDict, "zeroNumeric"), nil))
    Assert(IsEqual(exchange.SafeNumberOmitZero(inputDict, "zeroString"), nil))
    Assert(IsEqual(exchange.SafeNumberOmitZero(inputDict, "undefined"), nil))
    Assert(IsEqual(exchange.SafeNumberOmitZero(inputDict, "emptyString"), nil))
    Assert(!IsEqual(exchange.SafeNumberOmitZero(inputDict, "floatNumeric"), nil))
    Assert(!IsEqual(exchange.SafeNumberOmitZero(inputDict, "floatString"), nil))
}