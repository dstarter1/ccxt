package ccxt

type Myokx struct {
   *myokx
   Core *myokx
}

func NewMyokx(userConfig map[string]interface{}) Myokx {
   p := &myokx{}
   p.Init(userConfig)
   return Myokx{
       myokx: p,
       Core:  p,
   }
}

// PLEASE DO NOT EDIT THIS FILE, IT IS GENERATED AND WILL BE OVERWRITTEN:
// https://github.com/ccxt/ccxt/blob/master/CONTRIBUTING.md#how-to-contribute-code


