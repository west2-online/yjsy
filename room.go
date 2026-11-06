package yjsy

// 考场查询
//
// Deprecated: 请直接使用 GetSemesterCourses 并读取 Course.ExamTime。
func (s *Student) GetExamRoom(req ExamRoomReq) ([]*ExamRoomInfo, error) {
	courses, err := s.GetSemesterCourses(req.Term)
	if err != nil {
		return nil, err
	}

	examInfos := make([]*ExamRoomInfo, 0, len(courses))
	for _, course := range courses {
		if course == nil {
			continue
		}
		examInfo := &ExamRoomInfo{
			CourseName: course.Name,
			Credit:     course.Credits,
			Teacher:    course.Teacher,
		}
		if course.ExamTime != nil {
			examInfo.Date = course.ExamTime.Date
			examInfo.Time = course.ExamTime.Time
			examInfo.Location = course.ExamTime.Location
		}
		examInfos = append(examInfos, examInfo)
	}
	return examInfos, nil
}
