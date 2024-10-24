package crew_service_tests

//var testCrewServiceAttachParticipantToCrew = []struct {
//	testName  string
//	inputData struct {
//		participantID uuid.UUID
//		crewID        uuid.UUID
//		helmsman      int
//		active        int
//	}
//	prepare     func(fields *crewServiceFields)
//	checkOutput func(t *testing.T, crew *models.Crew, err error)
//}{
//	{
//		testName: "Success",
//		inputData: struct {
//			participantID uuid.UUID
//			crewID        uuid.UUID
//			helmsman      int
//			active        int
//		}{
//			uuid.New(),
//			uuid.New(),
//			1,
//			1,
//		},
//		prepare: func(fields *crewServiceFields) {
//			fields.crewRepoMock.EXPECT().AttachParticipantToCrew(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
//		},
//		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
//			assert.NoError(t, err)
//		},
//	},
//	{
//		testName: "attach participant to crew error",
//		inputData: struct {
//			participantID uuid.UUID
//			crewID        uuid.UUID
//			helmsman      int
//			active        int
//		}{
//			uuid.New(),
//			uuid.New(),
//			1,
//			1,
//		},
//		prepare: func(fields *crewServiceFields) {
//			fields.crewRepoMock.EXPECT().AttachParticipantToCrew(gomock.Any(), gomock.Any(), gomock.Any()).Return(repository_errors.UpdateError)
//		},
//		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
//			assert.Error(t, err)
//			assert.Equal(t, repository_errors.UpdateError, err)
//		},
//	},
//	{
//		testName: "input data error",
//		inputData: struct {
//			participantID uuid.UUID
//			crewID        uuid.UUID
//			helmsman      int
//			active        int
//		}{
//			uuid.New(),
//			uuid.New(),
//			-1,
//			-1,
//		},
//		prepare: func(fields *crewServiceFields) {
//			fields.crewRepoMock.EXPECT().AttachParticipantToCrew(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
//		},
//		checkOutput: func(t *testing.T, crew *models.Crew, err error) {
//			assert.NoError(t, err)
//		},
//	},
//}
//
//func TestCrewServiceAttachParticipantToCrew(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	fields := initCrewServiceFields(ctrl)
//	crewService := initCrewService(fields)
//
//	for _, tt := range testCrewServiceAttachParticipantToCrew {
//		t.Run(tt.testName, func(t *testing.T) {
//			tt.prepare(fields)
//			err := crewService.AttachParticipantToCrew(tt.inputData.participantID, tt.inputData.crewID, tt.inputData.helmsman)
//			tt.checkOutput(t, nil, err)
//		})
//	}
//}
