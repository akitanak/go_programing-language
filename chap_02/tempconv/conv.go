package tempconv

// CToFは摂氏の温度を華氏に変換します。
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToCは華氏の温度を摂氏に変換します。
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToKは摂氏を絶対温度に変換します。
func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

// KToCは絶対温度を摂氏に変換します。
func KToC(k Kelvin) Celsius { return Celsius(k) + AbsoluteZeroC }
