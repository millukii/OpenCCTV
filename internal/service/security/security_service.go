package security

/*
SecurityService

Management and Configuration
• Get or set advanced parameters of security management
Request URL: GET or PUT /ISAPI/Security/advanced?format=json
• Get or set security questions in batch
Request URL: GET or PUT /ISAPI/Security/questionConfiguration
• Search for security logs
Request URL: POST /ISAPI/ContentMgmt/security/logSearch
• Security Mode Level of Private Protocol
• Get configuration capability of security mode level of private protocol
Request URL: GET /ISAPI/Security/CommuMode/capabilities?format=json
• Get parameters of security mode level of private protocol
Request URL: GET /ISAPI/Security/CommuMode?format=json
• Set parameters of security mode level of private protocol
Request URL: PUT /ISAPI/Security/CommuMode?format=json
• Get capability of selecting certificate
Request URL: GET /ISAPI/Security/certificate/select/capabilities?format=json
• Get or set parameters of selecting certificate
Request URL: GET or PUT /ISAPI/Security/certificate/select/<functinName>?format=json

Network Certificate
• Get CA (Certificate Authority) certificate capability
Request URL: GET /ISAPI/Security/deviceCertificate/capabilities?format=json
• Import network certificate to device
Request URL: PUT /ISAPI/Security/deviceCertificate
• Import CA (Certificate Authority) certificate to device
Request URL: PUT /ISAPI/Security/deviceCertificate?customID=
• Network Certificate Search
• Get capability of certificate search
Request URL: GET /ISAPI/Security/deviceCertificate/certificates/capabilities?format=json
• Search for certificate information in a batch
Request URL: GET /ISAPI/Security/deviceCertificate/certificates?format=json
• Search for information of a specific device certificate
Request URL: GET /ISAPI/Security/deviceCertificate/certificates/<customID>?format=json
• Delete network certificate
Request URL: DELETE /ISAPI/Security/deviceCertificate/certificates/<ID> or /ISAPI/Security/
deviceCertificate/certificates/<customID>?format=json

Authentication Certificate
• Get, generate, or delete certificate signature request
Request URL: GET, PUT , or DELETE /ISAPI/Security/serverCertificate/certSignReq
• Download authentication certificate
Request URL: GET /ISAPI/Security/serverCertificate/downloadCertSignReq
• Get, install, or delete authentication certificate
Request URL: GET, PUT , or DELETE /ISAPI/Security/serverCertificate/certificate
• Authentication Certificate Status Search
• Get status of one authentication certificate
Request URL: GET /ISAPI/Security/deviceCertificate/certificates/<ID>/status?format=json
• Get status of all authentication certificates
Request URL: GET /ISAPI/Security/deviceCertificate/certificates/status?format=json
• Generate one authentication certificate
Request URL: PUT /ISAPI/Security/deviceCertificate/certificates/<ID>/recreate?format=json
• Generate all authentication certificates
Request URL: PUT /ISAPI/Security/deviceCertificate/certificates/recreate?format=json
Client/Server Certificate
• Get client/server certificate capability
Request URL: GET /ISAPI/Security/serverCertificate/capabilities?format=json
• Generate PKCS#10 signature request of client/server certificate
Request URL: POST /ISAPI/Security/serverCertificate/certSignReq?customID=
• Get or generate PKCS#10 signature request of client/server self-signed certificate
Request URL: GET or PUT /ISAPI/Security/serverCertificate/selfSignCert?customID=
• Get information of multiple client/server certificates in a batch
Request URL: GET /ISAPI/Security/serverCertificate/certificates?format=json
• Get or delete information of a specific client/server certificate
Request URL: GET or DELETE /ISAPI/Security/serverCertificate/certificates/<customID>?
format=json
• Import client/server certificate to device
Request URL: POST /ISAPI/Security/serverCertificate/certificate?customID=
• Export client/server certificate
Request URL: GET /ISAPI/Security/serverCertificate/downloadCertSignReq?customID=
Login by digest
Request URL: GET /ISAPI/Security/userCheck
• Lock for illegal login
Request URL: GET or PUT /ISAPI/Security/illegalLoginLock
• Get or set maximum failed login attempts
Request URL: GET or PUT /ISAPI/Security/loginLinkNum?format=json

*/
type SecurityService interface {
	GetSecurityParameters()
	GetSecurityQuestions()
	ResetPassword(recoverymail string)
}

type securityService struct{}

func NewSecurityService() SecurityService {
	return &securityService{}
}

/*
A recovery email is added or configured for resetting the password as required. The admin user can
set a recovery email after activating the device, and then receive the verification code from the
manufacturer via the recovery email to reset the device password.


*/
func (svc *securityService) ResetPassword(recoverymail string) {
	/*	Check if the device supports setting recovery email by the request URL: GET /ISAPI/Security/
			capabilities .

			The recovery email configuration capability is returned in the message JSON_SecurityEmailCap
		by the output parameter pointer (lpOutputParam).
		Optional: Get the existing or configured recovery email parameters for reference by the request
		URL: GET /ISAPI/Security/email/parameter?format=json .
		Set the recovery email for the device by the request URL: PUT /ISAPI/Security/email/
		parameter?format=json .
		Get the QR code of the recovery email by the request URL: GET /ISAPI/Security/email/qrCode?
		format=json .
		Send the QR code to manufacturer service email address by any other emails for verification.
		Receive the verification code in the configured recovery email.
		Enter the received verification code for resetting the device password by the request URL: PUT /
		ISAPI/Security/emailCertification?format=json .

	*/
}

func (svc *securityService) GetSecurityQuestions() {}

func (svc *securityService) GetSecurityParameters() {

}
