copy participants(id, name, category, gender, birthdate, coach_name) from 'E:\PPO_BMSTU\db\data\participants_data.csv' delimiter ';' csv header;
copy ratings(id, name, class, blowout_cnt) from 'E:\PPO_BMSTU\db\data\ratings_data.csv' delimiter ';' csv header;
copy crews(id, rating_id, class, sail_num ) from 'E:\PPO_BMSTU\db\data\crews_data.csv' delimiter ';' csv header;
copy races(id, rating_id, number, class, date) from 'E:\PPO_BMSTU\db\data\races_data.csv' delimiter ';' csv header;
copy judges( id, name, login, password, role, post) from 'E:\PPO_BMSTU\db\data\judges_data.csv' delimiter ';' csv header;
copy protests(id, race_id, rating_id, judge_id, rule_num, review_date, status, comment) from 'E:\PPO_BMSTU\db\data\protests_data.csv' delimiter ';' csv header;
copy crew_protest(id, crew_id, protest_id, crew_status) from 'E:\PPO_BMSTU\db\data\crew_protest_data.csv' delimiter ';' csv header;
copy crew_race(id, crew_id, race_id, points, spec_circumstance) from 'E:\PPO_BMSTU\db\data\crew_race_data.csv' delimiter ';' csv header;
copy participant_crew(id, participant_id, crew_id, helmsman, active) from 'E:\PPO_BMSTU\db\data\participant_crew_data.csv' delimiter ';' csv header;
copy judge_rating(id, judge_id, rating_id) from 'E:\PPO_BMSTU\db\data\judge_rating_data.csv' delimiter ';' csv header;