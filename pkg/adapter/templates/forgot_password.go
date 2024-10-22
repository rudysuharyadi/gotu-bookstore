package templates

type ForgotPasswordData struct {
	Email   string
	Token   string
	BaseURL string
}

// TODO: work with templates here,
// there should be another 1 page.
// So, email only link,
// then link will open html,
// then user input password and new password, submit.
// token and email can be pass it over.
var ForgotPasswordTemplates = `
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Registration Form</title>
<style>
    label {
        display: block;
        margin-bottom: 5px;
    }
</style>
</head>
<body>

<form action="{{.BaseURL}}/main-service/v1/auth/forgot-password/confirmation" method="post">
    <input type="hidden" id="email" name="email" value="{{.Email}}">
    <input type="hidden" id="token" name="token" value="{{.Token}}">

    <label for="password">Password:</label>
    <input type="password" id="password" name="password" required>

    <label for="confirm_password">Confirm Password:</label>
    <input type="password" id="confirm_password" name="confirm_password" required>

    <input type="submit" value="Submit">
</form>

</body>
</html>
`
