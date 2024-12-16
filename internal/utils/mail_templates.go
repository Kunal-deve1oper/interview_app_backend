package utils

import "fmt"

func SubmitMailTemplate(candidateName, jobName, platformName, companyName, candidateEmail, supportEmail string) string {
	template := fmt.Sprintf(`
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>Application Confirmation</title>
			</head>
			<body style="font-family: Arial, sans-serif; line-height: 1.6; margin: 0; padding: 0; background-color: #f4f4f9;">
				<table width="100%%" border="0" cellspacing="0" cellpadding="20" style="max-width: 600px; margin: 20px auto; background: #ffffff; border: 1px solid #ddd; border-radius: 5px;">
					<tr>
						<td style="text-align: center; background: #4CAF50; color: #ffffff; padding: 10px 20px; border-radius: 5px 5px 0 0;">
							<h1 style="margin: 0;">Application Received</h1>
						</td>
					</tr>
					<tr>
						<td style="padding: 20px;">
							<p>Dear %s,</p>
							<p>Thank you for applying for the <strong>%s</strong> position at <strong>%s</strong> through <strong>%s</strong>. We have successfully received your application.</p>
							<h3 style="margin-bottom: 10px;">Application Details:</h3>
							<ul style="list-style: none; padding: 0;">
								<li><strong>Position Applied For:</strong> %s</li>
								<li><strong>Company Name:</strong> %s</li>
								<li><strong>Email Address:</strong> <a href="mailto:%s">%s</a></li>
							</ul>
							<p>We've forwarded your application to <strong>%s</strong>. Their recruitment team will review your submission, and they'll reach out to you directly if your profile matches their requirements.</p>
							<p>If you have any questions about the process, feel free to contact us at <a href="mailto:%s">%s</a>.</p>
							<p>Thank you for using <strong>%s</strong>. We wish you the best of luck with your application!</p>
							<p>Warm regards,</p>
							<p><strong>%s</strong><br>
							%s<br>
							%s</p>
						</td>
					</tr>
					<tr>
						<td style="text-align: center; background: #f4f4f9; color: #888; font-size: 12px; padding: 10px; border-radius: 0 0 5px 5px;">
							&copy; 2024 %s. All rights reserved.
						</td>
					</tr>
				</table>
			</body>
			</html>
	`, candidateName, jobName, companyName, platformName, jobName, companyName, candidateEmail, candidateEmail, companyName, supportEmail, supportEmail, platformName, candidateName, jobName, platformName, platformName)
	return template
}
