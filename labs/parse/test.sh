
curl -X POST \
-H "X-Parse-Application-Id: my-app-id" \
-H "Content-Type: application/json" \
-d '{"score":1337,"playerName":"Sean Plott","cheatMode":false}' \
http://localhost:1337/parse/classes/GameScore

curl -X GET \
  -H "X-Parse-Application-Id: my-app-id" \
  http://localhost:1337/parse/classes/GameScore/VOcGD7sMG4