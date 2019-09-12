# internet-tester

On first time you run the program it will create a settings.json in the same directory the program is in. Edit that file so the program knows how it should work
```json
{
	"debug":false,
	"sec_to_wait":30,
	"url_to_test":"http://SOME_GOOD_URL/"
}
```

Install the program as a service by running a command prompt in Administrator mode and run the following
```
internet-tester.exe -service install
```
It will log to the windows application log
