/* 1. Выборки пользователей, у которых количество постов больше, чем у пользователя их пригласившего. */
SELECT u1.*
FROM users u1
JOIN users u2 ON u1.invited_by_user_id = u2.id
WHERE u1.posts_qty > u2.posts_qty;

/*---------------------result--------------------------------------

id | group_id | invited_by_user_id |      name      | posts_qty 
----+----------+--------------------+----------------+-----------
  2 |        1 |                  1 | Пользователь 2 |         2
  3 |        1 |                  2 | Пользователь 3 |         5
  4 |        2 |                  3 | Пользователь 4 |         7

------------------------------------------------------------------*/

/* 2. Выборки пользователей, имеющих максимальное количество постов в своей группе. */
SELECT u.*
FROM users u
JOIN (
    SELECT group_id, MAX(posts_qty) as max_posts
    FROM users
    GROUP BY group_id
) as grp_max ON u.group_id = grp_max.group_id AND u.posts_qty = grp_max.max_posts;

/*----------------------result------------------------------------

id | group_id | invited_by_user_id |      name      | posts_qty 
----+----------+--------------------+----------------+-----------
  3 |        1 |                  2 | Пользователь 3 |         5
  4 |        2 |                  3 | Пользователь 4 |         7

-------------------------------------------------------------------*/

/* 3. Выборки групп, количество пользователей в которых превышает 10000. */
SELECT g.id AS group_id, g.name AS group_name, COUNT(u.id) AS user_count
FROM groups g
JOIN users u ON g.id = u.group_id
GROUP BY g.id, g.name
HAVING COUNT(u.id) > 10000;

/*----------------------result------------------------------------

group_id | group_name | user_count 
----------+------------+------------
(0 rows)
В таблице users нет столько пользователей. 

-------------------------------------------------------------------*/

/* 4. Выборки пользователей, у которых пригласивший их пользователь из другой группы. */
SELECT u1.*
FROM users u1
JOIN users u2 ON u1.invited_by_user_id = u2.id
WHERE u1.group_id <> u2.group_id;

/*----------------------result------------------------------------

id | group_id | invited_by_user_id |      name      | posts_qty 
----+----------+--------------------+----------------+-----------
  4 |        2 |                  3 | Пользователь 4 |         7

-------------------------------------------------------------------*/

/* 5. Выборки групп с максимальным количеством постов у пользователей. */
WITH GroupPosts AS (
    SELECT group_id, SUM(posts_qty) AS total_posts
    FROM users
    GROUP BY group_id
)

SELECT g.*
FROM groups g
JOIN GroupPosts gp ON g.id = gp.group_id
WHERE gp.total_posts = (SELECT MAX(total_posts) FROM GroupPosts);

/*----------------------result------------------------------------

id |   name   
----+----------
  2 | Группа 2

-------------------------------------------------------------------*/