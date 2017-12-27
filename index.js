function main(ARGUMENTS) {
    const EXPECTED_QUANTITY_ARGUMENTS = 0;
    const ACTUAL_QUANTITY_ARGUMENTS = ARGUMENTS.length;

    if (ACTUAL_QUANTITY_ARGUMENTS === EXPECTED_QUANTITY_ARGUMENTS) {
        if (allArgumentsAreValid(ARGUMENTS)) {
            const TWILIO = require("twilio");

            const TWILIO_ACCOUNT_SID = process.env.TWILIO_ACCOUNT_SID;
            const TWILIO_AUTH_TOKEN = process.env.TWILIO_AUTH_TOKEN;
            const SENDER_PHONE_NUMBER = process.env.SENDER_PHONE_NUMBER;
            const RECEIVER_PHONE_NUMBER = process.env.RECEIVER_PHONE_NUMBER;
            const MESSAGE_BODY = "Message Body";

            let client = new TWILIO(TWILIO_ACCOUNT_SID, TWILIO_AUTH_TOKEN);

            client.messages.create({
                to: RECEIVER_PHONE_NUMBER,
                from: SENDER_PHONE_NUMBER,
                body: MESSAGE_BODY
            });

            console.log("You SMS with message body, \"" + MESSAGE_BODY + "\", was sent from phone number, \"" + SENDER_PHONE_NUMBER + "\", to phone number, \"" + RECEIVER_PHONE_NUMBER + "\", successfully.");
        } else {
            // Error messages handled in the checking function (bad design choice).
        }
    } else {
        console.error("Wrong Quantity Arguments Error:");
        console.error("Expected:");
        console.error("--- " + EXPECTED_QUANTITY_ARGUMENTS);
        console.error("Actual:");
        console.error("--- " + ACTUAL_QUANTITY_ARGUMENTS + ".");
        console.error("Your arguments were:");
        console.error("--- " + JSON.stringify(ARGUMENTS));
        console.error("Argument reference:");
        console.error("--- Argument 1: Message Body");
    }
}

function allArgumentsAreValid(ARGUMENTS) {
    if (process.env.TWILIO_ACCOUNT_SID.length === 34) {
        if (process.env.TWILIO_AUTH_TOKEN.length === 32) {
            if (process.env.SENDER_PHONE_NUMBER.length === 11 && isAnIntParsableString(process.env.SENDER_PHONE_NUMBER)) {
                if (process.env.RECEIVER_PHONE_NUMBER.length === 11 && isAnIntParsableString(process.env.RECEIVER_PHONE_NUMBER)) {
                    return true;
                } else {
                    console.error("Invalid Environment Variable Error: Receiver phone number must be a valid U.S. domestic phone number (don't forget the \"1\" at the front of the number).");
                    return false;
                }
            } else {
                console.error("Invalid Environment Variable Error: Sender phone number must be a valid U.S. domestic phone number (don't forget the \"1\" at the front of the number).");
                return false;
            }
        } else {
            console.error("Invalid Environment Variable Error: Twilio Auth Token must be 32 characters long.");
            return false;
        }
    } else {
        console.error("Invalid Environment Variable Error: Twilio Account SID must be 34 characters long.");
        return false;
    }
}

function isAnIntParsableString(argument) {
    return !isNaN(parseInt(argument));
}

const ALL_ARGUMENTS = process.argv; // Fetches all items passed to node, argv[0] = location of node interpreter, argv[1] = location of JavaScript file being executed, and all else are the user-passed arguments.
const ARGUMENTS = ALL_ARGUMENTS.slice(2);

main(ARGUMENTS);