/*
   To upload through terminal you can use: curl -F "image=@firmware.bin" esp32-webupdate.local/update
 */

#include <WiFi.h>
#include <WiFiClient.h>


// my imports
#include <Adafruit_MPU6050.h>
#include <Adafruit_Sensor.h>
#include <Wire.h>
#include <ArduinoJson.h>

#ifndef STASSID
#define STASSID "Tenda"
#define STAPSK "inter100200"
#endif

const char* host = "esp32-webupdate";
const char* ssid = STASSID;
const char* password = STAPSK;


// my vars
const uint16_t port = 8090;
const char* pc = "192.168.1.108";
#define LED_BUILDIN 2
////mpu
Adafruit_MPU6050 mpu;

void setup(void) {

        Serial.begin(115200);
        Serial.println();
        Serial.println("Booting Sketch...");
        WiFi.mode(WIFI_AP_STA);
        WiFi.begin(ssid, password);

        while (WiFi.waitForConnectResult() != WL_CONNECTED) {
                WiFi.begin(ssid, password);
                Serial.println("WiFi failed, retrying.");
        }
        /* my setup
         */
        pinMode(LED_BUILDIN, OUTPUT);

        digitalWrite(LED_BUILDIN, HIGH);

        if (!mpu.begin()) {
                while (1) {
                        delay(10);
                }
        }

        mpu.setAccelerometerRange(MPU6050_RANGE_8_G);
        mpu.setGyroRange(MPU6050_RANGE_500_DEG);
        mpu.setFilterBandwidth(MPU6050_BAND_21_HZ);
        delay(100);
}

void loop(void) {

        sensors_event_t a, g, temp;
        mpu.getEvent(&a, &g, &temp);


        StaticJsonDocument<192> doc;

        JsonObject mpu = doc.createNestedObject("mpu");

        doc["hall"] = hallRead();
        mpu["status"] = "true";

        JsonObject mpu_acceleration = mpu.createNestedObject("acceleration");
        mpu_acceleration["x"] = a.acceleration.x;
        mpu_acceleration["y"] = a.acceleration.y;
        mpu_acceleration["z"] = a.acceleration.z;

        JsonObject mpu_rotation = mpu.createNestedObject("rotation");
        mpu_rotation["x"] = g.gyro.x;
        mpu_rotation["y"] = g.gyro.y;
        mpu_rotation["z"] = g.gyro.z;

        mpu["temp"] = temp.temperature;

        String output;

        serializeJson(doc, output);

        WiFiClient client;

        if (!client.connect(pc, port)) {
                digitalWrite(LED_BUILDIN, HIGH);
                return;
        }

        client.print(output);
        client.stop();
        delay(20);
        digitalWrite(LED_BUILDIN, LOW);
}
