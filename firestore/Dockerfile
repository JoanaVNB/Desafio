FROM node:19-alpine
RUN apk update && apk add bash
RUN apk add openjdk11

RUN npm install -g firebase-tools

WORKDIR /app
COPY firebase.json .

ENV FIREBASE_PROJECT_ID=desafio-c0479

EXPOSE 9090

CMD ["sh", "-c", "firebase --project=desafio-c0479 emulators:start --only firestore" ]
