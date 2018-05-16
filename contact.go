package main

import (
	"net/smtp"

	"github.com/gin-gonic/gin"
)

// showContactForm executes the contact form
func showContactForm(c *gin.Context) {
	render(c, gin.H{
		"Title": "Contact Form",
	}, "contact.html")
}

// contactPost post the contact details and send to the given mail ID's
func contactPost(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	message := c.PostForm("message")
	var mailTemplateHTMLL = `
	<html>
		<head>
			<title>Email Template Sample</title>
			<meta http-equiv="Content-Type" content="text/html" charset="utf-8" />
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
		</head>
		<style>
			@media only screen and (min-width:320px) and (max-width: 599px) {
				body>table {
					width: 100% !important;
				}
			}
		</style>
		<body style="margin:0;font-family: arial,sans-serif;color:#000;background:#fff;">
			<table align="center" style="margin: 15px auto;width: 700px;padding:0px;cellpadding:0;cellspacing:0;border-spacing:0;outline:0;border:1px solid #ec1c24;">
				<tbody>
					<tr>
						<td align="center" style="padding:15px 5px;cellpadding:0;cellspacing:0;border-spacing:0;outline:0;border-collapse: collapse;border-bottom:1px solid #ec1c24;width:100%;">
							<img src="../images/logo.png" alt="Velan Info Services" title="Velan Info Services" />
						</td>
					</tr>

					<tr>
						<td style="padding:10px;cellpadding:0;cellspacing:0;border-spacing:0;outline:0;border-collapse: collapse;">
							<p style="font-size: 16px;margin:0;">It is a long established fact that a reader will be distracted by the readable content of Content.It is a long established fact that a reader will be distracted by the readable content of Content.</p>
						</td>
					</tr>

					<tr>
						<td style="padding:5px 25px;cellpadding:0;cellspacing:0;border-spacing:0;outline:0;border-collapse: collapse;">
							<b style="font-size: 16px;margin:0;color:#000000;">Employee Details:</b>
						</td>
					</tr>

					<tr>
						<td style="padding:10px 25px;cellpadding:0;cellspacing:0;border-spacing:0;outline:0;background:#fff;border-collapse: collapse;">
							<table align="center" style="padding:0px;margin: 0;cellpadding:0;cellspacing:0;border-spacing:0;outline:0;width:100%;border:1px solid #ec1c24;border-bottom:none;">
								<tbody>
									<tr>
										<td style="padding:0 5px;cellpadding:0;cellspacing:0;border-spacing:0;outline:0;border-collapse: collapse;width:50%;">
											<ul style="margin:0;padding:0 0 0 15px;list-style: none;">
												<li style="font-size: 14px;margin:5px 0 5px 0;"> Name:</li>
											</ul>
										</td>
										<td style="padding:0 5px;cellpadding:0;cellspacing:0;border-spacing:0;outline:0;border-collapse: collapse;width:50%;border-left:1px solid #ec1c24;">
											<ul style="margin:0;padding:0 0 0 15px;list-style: none;">
												<li style="font-size: 14px;margin:5px 0 5px 0;">` + name + `</li>
											</ul>
										</td>
									</tr>
								</tbody>
							</table>

							<table align="center" style="padding:0px;margin: 0;cellpadding:0;cellspacing:0;border-spacing:0;outline:0;width:100%;border:1px solid #ec1c24;border-bottom:none;">
								<tbody>
									<tr>
										<td style="padding:0 5px;cellpadding:0;cellspacing:0;border-spacing:0;outline:0;border-collapse: collapse;width:50%;">
											<ul style="margin:0;padding:0 0 0 15px;list-style: none;">
												<li style="font-size: 14px;margin:5px 0 5px 0;">Email:</li>
											</ul>
										</td>
										<td style="padding:0 5px;cellpadding:0;cellspacing:0;border-spacing:0;outline:0;border-collapse: collapse;width:50%;border-left:1px solid #ec1c24;">
											<ul style="margin:0;padding:0 0 0 15px;list-style: none;">
												<li style="font-size: 14px;margin:5px 0 5px 0;">` + email + `</li>
											</ul>
										</td>
									</tr>
								</tbody>
							</table>

							<table align="center" style="padding:0px;margin: 0;cellpadding:0;cellspacing:0;border-spacing:0;outline:0;width:100%;border:1px solid #ec1c24;">
								<tbody>
									<tr>
										<td style="padding:0 5px;cellpadding:0;cellspacing:0;border-spacing:0;outline:0;border-collapse: collapse;width:50%;">
											<ul style="margin:0;padding:0 0 0 15px;list-style: none;">
												<li style="font-size: 14px;margin:5px 0 5px 0;">Phone:</li>
											</ul>
										</td>
										<td style="padding:0 5px;cellpadding:0;cellspacing:0;border-spacing:0;outline:0;border-collapse: collapse;width:50%;border-left:1px solid #ec1c24;">
											<ul style="margin:0;padding:0 0 0 15px;list-style: none;">
												<li style="font-size: 14px;margin:5px 0 5px 0;">` + phone + `</li>
											</ul>
										</td>
									</tr>
								</tbody>
							</table>

							<table align="center" style="padding:0px;margin: 0;cellpadding:0;cellspacing:0;border-spacing:0;outline:0;width:100%;border:1px solid #ec1c24;">
								<tbody>
									<tr>
										<td style="padding:0 5px;cellpadding:0;cellspacing:0;border-spacing:0;outline:0;border-collapse: collapse;width:50%;">
											<ul style="margin:0;padding:0 0 0 15px;list-style: none;">
												<li style="font-size: 14px;margin:5px 0 5px 0;">Message:</li>
											</ul>
										</td>
										<td style="padding:0 5px;cellpadding:0;cellspacing:0;border-spacing:0;outline:0;border-collapse: collapse;width:50%;border-left:1px solid #ec1c24;">
											<ul style="margin:0;padding:0 0 0 15px;list-style: none;">
												<li style="font-size: 14px;margin:5px 0 5px 0;">` + message + `</li>
											</ul>
										</td>
									</tr>
								</tbody>
							</table>
						</td>
					</tr>

					<tr>
						<td style="padding:10px;cellpadding:0;cellspacing:0;border-spacing:0;outline:0;border-collapse: collapse;">
							<p style="font-size: 16px;margin:0;">It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout. The point of using Lorem Ipsum is that it has a more-or-less normal distribution of letters, as opposed to
							using 'Content here, content here', making it look like readable English. Many desktop publishing packages and web page editors now use Lorem Ipsum as their default model text, and a search for 'lorem ipsum' will uncover many web
							sites still in their infancy. Various versions have evolved over the years, sometimes by accident, sometimes on purpose (injected humour and the like).</p>
						</td>
					</tr>

					<tr>
						<td align="center" style="padding:5px 25px;cellpadding:0;cellspacing:0;border-spacing:0;outline:0;background:#ec1c24;border-collapse: collapse;">
							<p style="font-size: 12px;margin:0;color:#fff;">&copy; Golang Email Template.</p>
						</td>
					</tr>
				</tbody>
			</table>
		</body>
	</html>`

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := "To:" + email + "\r\nSubject: Test Mail \n" + mime + "\r\n\r\n" + mailTemplateHTMLL
	auth := smtp.PlainAuth("Viknesh NM", "ultimateviknesh@gmail.com", "vickyvicky23", "smtp.gmail.com")
	err = smtp.SendMail("smtp.gmail.com:587", auth, "Viknesh NM", []string{"ultimateviknesh@gmail.com", "viknesh.dev@velaninfo.com"}, []byte(body))
	if err != nil {
		panic(err)
	}
	render(c, gin.H{
		"Success": "Success",
	}, "contact.html")
}
