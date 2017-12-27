# Crypto Price Fetcher

### Prerequisites:
You will need node and npm. At the time of writing my versions are `node: v8.1.3` and `npm: 5.6.0`.

Install the dependencies for the project by running `npm install`.

### Running:
Execute the following script to start the project with the proper environment variables. Note, all values surrounded by `~`, including the `~`'s themselves, should be replaced by an actual value.

`TWILIO_ACCOUNT_SID=~Your Twilio account SID~ TWILIO_AUTH_TOKEN=~Your Twilio auth token~ SENDER_PHONE_NUMBER=~full sender phone number~ RECEIVER_PHONE_NUMBER=~full receiver phone number~ node index.js`
