package tests

func (suite *e2eTestSuite) TestUserPasswordChange() {
	// Given: A user with email and password exists and is logged in
	suite.Run("Create and login user", func() {
		var err error
		email := "testuser@example.com"
		password := "password123"
		suite.judgeID, err = suite.aJudgeExistsWithEmailAndPassword(email, password)
		suite.NoError(err, "User creation should succeed")

		err = suite.theUserLogsInWithEmailAndPassword(email, password)
		suite.NoError(err, "User login should succeed")
	})

	// When: The user changes their password
	suite.Run("Change password", func() {
		newPassword := "newpassword123"
		err := suite.theUserChangesTheirPasswordTo(newPassword, suite.judgeID)
		suite.NoError(err, "Password change should succeed")
	})

	// Then: The password is changed
	suite.Run("Verify password change", func() {
		// Попробуйте войти с новым паролем
		newPassword := "newpassword123"
		err := suite.theUserLogsInWithEmailAndPassword("testuser@example.com", newPassword)
		suite.NoError(err, "Login with new password should succeed")
	})
}
