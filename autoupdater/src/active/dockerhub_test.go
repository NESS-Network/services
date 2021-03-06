package active

import (

"fmt"
"net/http"
"net/http/httptest"
"testing"
"time"

"github.com/skycoin/services/autoupdater/src/logger"
"github.com/stretchr/testify/mock"

)

const mockTokenResponse = `{
    "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsIng1YyI6WyJNSUlDK2pDQ0FwK2dBd0lCQWdJQkFEQUtCZ2dxaGtqT1BRUURBakJHTVVRd1FnWURWUVFERXpzeVYwNVpPbFZMUzFJNlJFMUVVanBTU1U5Rk9reEhOa0U2UTFWWVZEcE5SbFZNT2tZelNFVTZOVkF5VlRwTFNqTkdPa05CTmxrNlNrbEVVVEFlRncweE9EQXlNVFF5TXpBMk5EZGFGdzB4T1RBeU1UUXlNekEyTkRkYU1FWXhSREJDQmdOVkJBTVRPMVpCUTFZNk5VNWFNenBNTkZSWk9sQlFTbGc2VWsxQlZEcEdWalpQT2xZMU1sTTZRa2szV2pwU1REVk9PbGhXVDBJNlFsTmFSanBHVTFRMk1JSUJJakFOQmdrcWhraUc5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0NBUUVBMGtyTmgyZWxESnVvYjVERWd5Wi9oZ3l1ZlpxNHo0OXdvNStGRnFRK3VPTGNCMDRyc3N4cnVNdm1aSzJZQ0RSRVRERU9xNW5keEVMMHNaTE51UXRMSlNRdFY1YUhlY2dQVFRkeVJHUTl2aURPWGlqNFBocE40R0N0eFV6YTNKWlNDZC9qbm1YbmtUeDViOElUWXBCZzg2TGNUdmMyRFVUV2tHNy91UThrVjVPNFFxNlZKY05TUWRId1B2Mmp4YWRZa3hBMnhaaWNvRFNFQlpjWGRneUFCRWI2YkRnUzV3QjdtYjRRVXBuM3FXRnRqdCttKzBsdDZOR3hvenNOSFJHd3EwakpqNWtZbWFnWHpEQm5NQ3l5eDFBWFpkMHBNaUlPSjhsaDhRQ09GMStsMkVuV1U1K0thaTZKYVNEOFZJc2VrRzB3YXd4T1dER3U0YzYreE1XYUx3SURBUUFCbzRHeU1JR3ZNQTRHQTFVZER3RUIvd1FFQXdJSGdEQVBCZ05WSFNVRUNEQUdCZ1JWSFNVQU1FUUdBMVVkRGdROUJEdFdRVU5XT2pWT1dqTTZURFJVV1RwUVVFcFlPbEpOUVZRNlJsWTJUenBXTlRKVE9rSkpOMW82VWt3MVRqcFlWazlDT2tKVFdrWTZSbE5VTmpCR0JnTlZIU01FUHpBOWdEc3lWMDVaT2xWTFMxSTZSRTFFVWpwU1NVOUZPa3hITmtFNlExVllWRHBOUmxWTU9rWXpTRVU2TlZBeVZUcExTak5HT2tOQk5sazZTa2xFVVRBS0JnZ3Foa2pPUFFRREFnTkpBREJHQWlFQWdZTWF3Si9uMXM0dDlva0VhRjh2aGVkeURzbERObWNyTHNRNldmWTFmRTRDSVFEbzNWazJXcndiSjNmU1dwZEVjT3hNazZ1ZEFwK2c1Nkd6TjlRSGFNeVZ1QT09Il19.eyJhY2Nlc3MiOlt7InR5cGUiOiJyZXBvc2l0b3J5IiwibmFtZSI6ImxpYnJhcnkvbWFyaWFkYiIsImFjdGlvbnMiOlsicHVsbCJdfV0sImF1ZCI6InJlZ2lzdHJ5LmRvY2tlci5pbyIsImV4cCI6MTUzMTI4ODA5MCwiaWF0IjoxNTMxMjg3NzkwLCJpc3MiOiJhdXRoLmRvY2tlci5pbyIsImp0aSI6IkVneGgySENXUl94S0N3Q2NhQ0tCIiwibmJmIjoxNTMxMjg3NDkwLCJzdWIiOiIifQ.tc9TfkPlZDk_UF2eTOzUWz2uegRPgPs3fCGN2VGkI98yQbY0UwKpNpONdaHEy1BDJip5wJ7ff1f7iVqhgNfNtva64wMO3eKokZLViPb2HIcIxcWNUgaCIOMVr_RKxoD92PmCTFtHpMYqyXtjJgeZh437jJJJfOTj6QEeupanPSDjYWfDDFxxIvwSOd0CPbefxTHQZV3S7tNmNt3OHODIN-rBkUxFYgVO_7BuMVigqoQ29OoKeriPFGSxUKtQGJhs1b9KF7_QeZF8f71kNBVPNkSfetKSrTA4VJxPlS-JUjcpjwIB4t00LCRWnqYCveeYvFYfdXGwnscfIh2luEJAnw",
    "access_token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsIng1YyI6WyJNSUlDK2pDQ0FwK2dBd0lCQWdJQkFEQUtCZ2dxaGtqT1BRUURBakJHTVVRd1FnWURWUVFERXpzeVYwNVpPbFZMUzFJNlJFMUVVanBTU1U5Rk9reEhOa0U2UTFWWVZEcE5SbFZNT2tZelNFVTZOVkF5VlRwTFNqTkdPa05CTmxrNlNrbEVVVEFlRncweE9EQXlNVFF5TXpBMk5EZGFGdzB4T1RBeU1UUXlNekEyTkRkYU1FWXhSREJDQmdOVkJBTVRPMVpCUTFZNk5VNWFNenBNTkZSWk9sQlFTbGc2VWsxQlZEcEdWalpQT2xZMU1sTTZRa2szV2pwU1REVk9PbGhXVDBJNlFsTmFSanBHVTFRMk1JSUJJakFOQmdrcWhraUc5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0NBUUVBMGtyTmgyZWxESnVvYjVERWd5Wi9oZ3l1ZlpxNHo0OXdvNStGRnFRK3VPTGNCMDRyc3N4cnVNdm1aSzJZQ0RSRVRERU9xNW5keEVMMHNaTE51UXRMSlNRdFY1YUhlY2dQVFRkeVJHUTl2aURPWGlqNFBocE40R0N0eFV6YTNKWlNDZC9qbm1YbmtUeDViOElUWXBCZzg2TGNUdmMyRFVUV2tHNy91UThrVjVPNFFxNlZKY05TUWRId1B2Mmp4YWRZa3hBMnhaaWNvRFNFQlpjWGRneUFCRWI2YkRnUzV3QjdtYjRRVXBuM3FXRnRqdCttKzBsdDZOR3hvenNOSFJHd3EwakpqNWtZbWFnWHpEQm5NQ3l5eDFBWFpkMHBNaUlPSjhsaDhRQ09GMStsMkVuV1U1K0thaTZKYVNEOFZJc2VrRzB3YXd4T1dER3U0YzYreE1XYUx3SURBUUFCbzRHeU1JR3ZNQTRHQTFVZER3RUIvd1FFQXdJSGdEQVBCZ05WSFNVRUNEQUdCZ1JWSFNVQU1FUUdBMVVkRGdROUJEdFdRVU5XT2pWT1dqTTZURFJVV1RwUVVFcFlPbEpOUVZRNlJsWTJUenBXTlRKVE9rSkpOMW82VWt3MVRqcFlWazlDT2tKVFdrWTZSbE5VTmpCR0JnTlZIU01FUHpBOWdEc3lWMDVaT2xWTFMxSTZSRTFFVWpwU1NVOUZPa3hITmtFNlExVllWRHBOUmxWTU9rWXpTRVU2TlZBeVZUcExTak5HT2tOQk5sazZTa2xFVVRBS0JnZ3Foa2pPUFFRREFnTkpBREJHQWlFQWdZTWF3Si9uMXM0dDlva0VhRjh2aGVkeURzbERObWNyTHNRNldmWTFmRTRDSVFEbzNWazJXcndiSjNmU1dwZEVjT3hNazZ1ZEFwK2c1Nkd6TjlRSGFNeVZ1QT09Il19.eyJhY2Nlc3MiOlt7InR5cGUiOiJyZXBvc2l0b3J5IiwibmFtZSI6ImxpYnJhcnkvbWFyaWFkYiIsImFjdGlvbnMiOlsicHVsbCJdfV0sImF1ZCI6InJlZ2lzdHJ5LmRvY2tlci5pbyIsImV4cCI6MTUzMTI4ODA5MCwiaWF0IjoxNTMxMjg3NzkwLCJpc3MiOiJhdXRoLmRvY2tlci5pbyIsImp0aSI6IkVneGgySENXUl94S0N3Q2NhQ0tCIiwibmJmIjoxNTMxMjg3NDkwLCJzdWIiOiIifQ.tc9TfkPlZDk_UF2eTOzUWz2uegRPgPs3fCGN2VGkI98yQbY0UwKpNpONdaHEy1BDJip5wJ7ff1f7iVqhgNfNtva64wMO3eKokZLViPb2HIcIxcWNUgaCIOMVr_RKxoD92PmCTFtHpMYqyXtjJgeZh437jJJJfOTj6QEeupanPSDjYWfDDFxxIvwSOd0CPbefxTHQZV3S7tNmNt3OHODIN-rBkUxFYgVO_7BuMVigqoQ29OoKeriPFGSxUKtQGJhs1b9KF7_QeZF8f71kNBVPNkSfetKSrTA4VJxPlS-JUjcpjwIB4t00LCRWnqYCveeYvFYfdXGwnscfIh2luEJAnw",
    "expires_in": 300,
    "issued_at": "2018-07-11T05:43:10.518662256Z"
}`

const mockRepositoryResponse = `{
    "schemaVersion": 2,
    "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
    "config": {
        "mediaType": "application/vnd.docker.container.image.v1+json",
        "size": 7619,
        "digest": "sha256:520fc647a087d0e055bcf411b8b196de3e31ef78a8596b5b78e078825b2072bb"
    }
}`

func TestToken(t *testing.T) {
	// Arrange
	updaterMock, dockerhubFetcher, tokenIssuer, repository := arrange()
	defer tokenIssuer.Close()
	defer repository.Close()

	// Action
	go dockerhubFetcher.Start()

	// Assert
	time.Sleep(time.Second * 2)
	updaterMock.AssertExpectations(t)
	dockerhubFetcher.Stop()
}

func arrange() (UpdaterMock, Fetcher, *httptest.Server, *httptest.Server) {
	// Mock for the Updater service, so it does not try to contact swarm
	updaterMock := &UpdaterMock{}

	// Mocks for the token issuer server and the dockerhub server
	tokenIssuer := httptest.NewServer(http.HandlerFunc(mockTokenIssuer))
	repository := httptest.NewServer(http.HandlerFunc(mockDockerRepository))

	// Logger and set mock response
	log := logger.NewLogger("service")
	updaterMock.On("Update", "service", "test/service:latest", log).Return(nil)

	// Create a dockerhub fetcher instance and setup server mocks
	dockerhubFetcher := newDockerHub(updaterMock,"/test/service",
		"latest","service","digest",log )

	dockerhubFetcher.TokenTemplate = tokenIssuer.URL + "/%s"
	dockerhubFetcher.Url = repository.URL
	dockerhubFetcher.SetInterval(time.Second * 1)

	return *updaterMock, dockerhubFetcher, tokenIssuer, repository
}

type UpdaterMock struct {
	mock.Mock
}

func (u *UpdaterMock) Update(service, version string, log *logger.Logger) chan error {
	u.Called(service, version, log)
	return make(chan error)
}

func mockTokenIssuer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, mockTokenResponse)
}

func mockDockerRepository(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, mockRepositoryResponse)
}
