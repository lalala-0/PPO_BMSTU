package tests

func (suite *e2eTestSuite) TestUserAuthenticationWith2FA() {
	// Given: A user with email and password exists
	suite.Run("Create user", func() {
		var err error
		email := "testuser@example.com"
		password := "password123"
		suite.judgeID, err = suite.aJudgeExistsWithEmailAndPassword(email, password)
		suite.NoError(err, "User creation should succeed")
	})

	// When: The user logs in, generates, and verifies the 2FA code
	suite.Run("Login and authenticate with 2FA", func() {
		email := "testuser@example.com"
		password := "password123"
		err := suite.theUserLogsInWithEmailAndPassword(email, password)
		suite.NoError(err, "User login should succeed")

		code, err := suite.theUserGeneratesA2FACode(suite.judgeID)
		suite.NoError(err, "2FA code generation should succeed")

		err = suite.theUserVerifiesThe2FACode(code, suite.judgeID)
		suite.NoError(err, "2FA code verification should succeed")
	})

	// Then: The user is authenticated
}
