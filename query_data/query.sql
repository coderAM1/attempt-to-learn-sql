SELECT r.*
FROM (
    SELECT u.userid, u.name, SUM(t.distancetraveled) as distancetravel, count(t.tripid) as trip_count
    FROM users AS u
    JOIN trips AS t on u.userid = t.userid
    GROUP BY u.userid
    ORDER BY distancetravel DESC
    LIMIT 100
) AS r ORDER BY r.name ASC;