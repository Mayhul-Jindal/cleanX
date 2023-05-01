# CleanX
An autonomous boat which can clean a sewage tank using the concept of floating plants. 

## How does it work ?
Using esp32 to connect different sensors and publishing the sensor data in realtime to our backend. Using the value of DO sensor to judge the speed of the boat and opensource tools to monitor the water quality index. Also using an AI model to predict when is the most of amount of waste generated in water and that data is used by out college to make effective decisions towards waste-water management system.   

[Prototype-1 of our boat](https://user-images.githubusercontent.com/95216160/235504780-4e90d7b4-433d-44ca-9494-c581c73eec67.webm)

## System design
Following is the schematics for our circuitry

<p float="left">
  <img src="https://user-images.githubusercontent.com/95216160/235494977-357f9bdc-4976-4f2b-bd5e-0d02e85297b9.png" width="51%" />
  <img src="https://user-images.githubusercontent.com/95216160/235506743-70a5e71a-4fa8-4c0d-bc7b-ea9c6af8d6c2.png" width="45%" /> 
</p>

Here is the system design which gave us the capability to monitor our boat in realtime

![image](https://user-images.githubusercontent.com/95216160/235504340-aca6add4-d5e5-45ec-9536-8f533d1323cc.png)
