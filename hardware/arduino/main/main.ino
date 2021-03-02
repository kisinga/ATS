/*
 * Copyright © 2021 Kisinga
 *  
 */

// include the GSM library
#include <MKRGSM.h>

//set up the buttons
const int lowestPin = 0;

const int highestPin = 10;

// Please enter your sensitive data in the Secret tab or arduino_secrets.h
// PIN Number

#define PINNUMBER ""

// initialize the library instances

GSM gsmAccess;

GSM_SMS sms;

// Array to hold the number a SMS is retreived from
char senderNumber[20];

void setup()
{
  // set pins 0 through 11 as outputs:
  for (int thisPin = lowestPin; thisPin <= highestPin; thisPin++)
  {
    pinMode(thisPin, OUTPUT);
  }

  // initialize serial communications and wait for port to open:

  Serial.begin(9600);

  while (!Serial)
  {

    ; // wait for serial port to connect. Needed for native USB port only
  }

  Serial.println("SMS Messages Receiver");

  // connection state

  bool connected = false;

  // Start GSM connection

  while (!connected)
  {

    if (gsmAccess.begin(PINNUMBER) == GSM_READY)
    {

      connected = true;
    }
    else
    {

      Serial.println("Not connected");

      delay(1000);
    }
  }
  Serial.println("GSM initialized");
  Serial.println("Waiting for messages");
}

// the loop function runs over and over again forever
void loop()
{
  //A loop for testing purposes only
  //   for (int thisPin = lowestPin; thisPin <= highestPin; thisPin++) {
  //          pressButton(thisPin);
  //    }
  // If there are any SMSs available()

  if (sms.available())
  {

    Serial.println("Message received from:");
    // Get remote number
    sms.remoteNumber(senderNumber, 20);

    Serial.println(senderNumber);

    // Discard all messages that dont start with #
    if (sms.peek() != '#')
    {
      Serial.println("Discarded SMS");
      sms.flush();
      return;
    }

    // Read message bytes and print them
    int c;

    while ((c = sms.read()) != -1)
    {
      int num = (char)c - '0';
      Serial.print(num);
      pressButton(num);
    }
    //press ok
    pressButton(10);
    Serial.println("\nEND OF MESSAGE");

    // Delete message from modem memory

    sms.flush();

    Serial.println("MESSAGE DELETED");
  }
}

void pressButton(int button)
{
  digitalWrite(button, HIGH); // Press the button
  delay(30);                  // wait for sometime for the CIU to register the keypress
  digitalWrite(button, LOW);  // unpress the button
  delay(50);                  // wait for sometime just in case there is a series of buttons to be pressed
}
