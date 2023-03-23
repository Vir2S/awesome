package database

func getAllApiKeysQuery() string {
	return "SELECT api_key FROM auth;"
}

func GetAllProfilesQuery() string {
	return "SELECT user.username, user_profile.user_id," +
		" user_profile.first_name, user_profile.last_name, user_profile.city, user_data.school" +
		" FROM user JOIN user_profile ON user.id = user_profile.user_id JOIN user_data ON user.id = user_data.user_id;"
}

func GetProfileQuery(username string) string {
	return "SELECT user.username, user_profile.user_id," +
		" user_profile.first_name, user_profile.last_name, user_profile.city, user_data.school" +
		" FROM user JOIN user_profile ON user.id = user_profile.user_id JOIN user_data ON user.id = user_data.user_id" +
		" WHERE user.username = '" + username + "';"
}
