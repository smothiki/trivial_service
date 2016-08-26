# trivial_service

This is a simple web services example that is used for teaching.

There are two web servers in this project.

www is a simple frontend web server that only responds to requests to '/'.

backend is a simple backend web server that usually returns the current time.  Usually because it has an
approximately 2% error rate.

www for it's part has some issues too. For one it generally never exceeds 10 requests per second.  Secondly
it crashes if the backend server returns an error.

Both web servers accept various command-line parameters and this should be pretty obvious from the code.
