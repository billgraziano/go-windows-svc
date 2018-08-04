# GO Windows Service Shell

This is based on the [GO Windows service example program](https://godoc.org/golang.org/x/sys/windows/svc/example) provided by the GO Project. 
It is a project shell to create a Windows service.

## Getting Started

The program compiles and runs on GO 1.8.  The generated executable accepts a single parameter.  The parameter values include:

* debug - runs the program from the command-line
* install - installs a windows service
* remove - removes the windows service
* start
* stop
* pause
* continue

## Installing and Updating a Service

After compiling an executable, the service can be installed from an Administrative command prompt.  Typing

    YourExecutable.EXE install 

will install the service.

To update the service, stop the service, replace the executble and restart the service.

The service can be removed from an Administrative command prompt by typing:

    YourExecutable.EXE remove 

## Customizing

The code exists in two packabages

* cmd/gosvc - Wrapper to control the service
* app - Your application

All service boilerplate code is in the four files in `cmd/gosvc` with a "svc_" prefix.  There should
be no need to modify this code.

The only code you should need to change is in main.go.

* `svcName` - This constant is the name of the installed service.  This is used for NET START and NET STOP commands.
* `svcNameLong` - This is the longer service name that appears in the Services control panel.
* `svcLauncher()` - This launches app.Run passing it a Windows Event Logger, the `svcName`, and the SHA1 hash from GIT.

You should also rename the `gosvc` directory to the name of your executable.

* `setup()` - This function is called to do any application setup.  If returns a non-nil error, the service will exit.
* `yourApp()` - This is launched as a GO routine.  This is the body of your application.  Here you can launch web servers, listeners, or whatever your application does.



## Advanced
### Logging
The `server` struct exposes a `winlog` variable that is a logger.  This will write to the console when running interactively and to the Winodws Application Event Log when running as a service.  I typically use this for any service errors, start and stop notification, and any issues reading configuration or setting up logging.

### Other 
This uses the [GO error wrapper package]("github.com/pkg/errors").  You can easily 
remove it if you prefer not to use it.










