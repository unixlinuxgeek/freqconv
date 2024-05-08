### Frequency Converter 

Run tests:
```
$ go test -v
=== RUN   TestHzToKHz
freqconv_test.go:15: Test TestHzToKHz Passed: 0.00869406303360217 kHz equal 0.00869406303360217 kHz
--- PASS: TestHzToKHz (0.00s)
=== RUN   TestHzToMHz
freqconv_test.go:28: Test TestHzToMHz Passed: 8.224785290419593e-06 MHz equal 8.224785290419593e-06 MHz
--- PASS: TestHzToMHz (0.00s)
=== RUN   TestHzToGHz
freqconv_test.go:41: Test TestHzToGHz Passed: 5.370736318423669e-09 GHz equal 5.370736318423669e-09 GHz
--- PASS: TestHzToGHz (0.00s)
=== RUN   TestKHzToHz
freqconv_test.go:54: Test TestKHzToHz Passed: 7710.221412824127 Hz equal 7710.221412824127 Hz
--- PASS: TestKHzToHz (0.00s)
=== RUN   TestKHzToMHz
freqconv_test.go:67: Test TestKHzToMHz Passed: 0.00806428697400276 MHz equal 0.00806428697400276 MHz
--- PASS: TestKHzToMHz (0.00s)
=== RUN   TestKhzToGhz
freqconv_test.go:80: Test TestKhzToGhz Passed: 1.5507884835699758e-06 GHz equal 1.5507884835699758e-06 GHz
--- PASS: TestKhzToGhz (0.00s)
=== RUN   TestMHzToHz
freqconv_test.go:92: Test TestMHzToHz Passed: 1.7255345856280653e+06 Hz equal 1.7255345856280653e+06 Hz
--- PASS: TestMHzToHz (0.00s)
=== RUN   TestMHzToKHz
freqconv_test.go:105: Test TestMHzToKHz Passed: 7922.914947210515 kHz equal 7922.914947210515 kHz
--- PASS: TestMHzToKHz (0.00s)
=== RUN   TestMHzToGHz
freqconv_test.go:117: Test TestMHzToGHz Passed: 0.0031991781706642627 GHz equal 0.0031991781706642627 GHz
--- PASS: TestMHzToGHz (0.00s)
PASS
ok      github.com/unixlinuxgeek/freqconv       0.002s
```
