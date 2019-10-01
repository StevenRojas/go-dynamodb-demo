# **DynamoDB - AWS SDK for Go**

The following is an explanation and a demo of how to use DynamoDB with Go. This document includes the following:

* [DynamoDB concepts](#dynamodb-concepts)
* [AWS SDK](#aws-sdk)
* [DynamoDB API](#dynamodb-api)

# **DynamoDB concepts**
## **What is Amazon DynamoDB**

 - It is a NoSQL database with high performance and scalability implemented using AWS SDK
 - Supports key-value and document data structures
 - Offers an encryption layer 
 - Is easy to store JSON documents in a DynamoDB table preserving JSON’s complexity and nested shape

## **DynamoDB core components**

 - **Tables:** DynamoDB stores data in tables
 - **Items:** It is a set of attributes. Each table contains zero or more items (similar to rows or records)
 - **Attributes:** Is a fundamental data element. (similar to columns)
 - **Primary keys**: Identify uniquely each item. There is two types:
	 - **Partition key:** It is used to generate a hash which defines the partition where the item will be stored
	 - **Partition key and sort key:** Also known as composite primary key, composed of two attributes. The items are stored in the corresponding partition and sorted order by sort key value. A composite primary key provides flexibility when querying data. **Primary Key** is also known as **hash attribute** and **Sort Key** is know as **range attribute**
 - **Secondary indexes:** Allows query the data using an alternate key, there could be one or more secondary indexes.
	 - An index belongs to a table
	 - DynamoDB handles automatically the indexes on each CRUD operation on the table
	 - When an index is created is possible to specify which attributes will be copied from the table to index table in order to have direct access to those attributes

![](https://lh4.googleusercontent.com/KPNBu-XSDL67Z-Y-scWptZS8_gcL4wzvGRZNdd14JNGQnnffgo1i7d_kTh3UaAQScQ3DHNEzWz9VKgtDvnRom8lK7t3wRYeNVhbTnBFxLNWASIEi14HifVdOh3dE0d9lai7-dnGxW3E)

 - **DynamoDB Streams:** Captures events related to data modification in a stream in near-real time and in the order that the events occurred
	 - Every event is stored as stream record
	 - If a new item is added to the table, the stream will store all of its attributes.
	 - If an item is updated, the stream captures the "before" and "after" modified attributes.
	 - If an item is deleted, the stream captures the entire item.
	 - Streams are stored for only 24 hours
	 - It is possible to use AWS Lambda in order to generate triggers

## ** Relational DB vs NoSQL DB**
![](https://lh5.googleusercontent.com/9JPgidu-fX9SOer5M6IBte7_DMsEyeaooyfWKgRu5nlH8q7PMAwXyH2CoZB-NG24WiXBx5ZcFyegP_QKoGZi6is1WC4-EY23FZ4xNsw3jReHKENqMiKvEB4OUjJfeayfm_S3jhfX63A)

![](https://lh5.googleusercontent.com/RZkMMDFrRneMLnf7ygcQ9OtY8uZN0_JtCvD_Q6S4IMzr3XhTCtTcBKVpKl4ZGW4SzgDrlfu20JYBLyycnrjuKF3rX2ByNpmNliQorwopBj49IE3omav1YIVsS5hCOij1TdCgHi09ZeE)

# **AWS SDK**
## **AWS SDK for Go**
The AWS SDK for Go provides APIs and utilities that developers can use to build Go applications that use AWS services:
 - Amazon Elastic Compute Cloud (Amazon EC2)
 - Amazon Simple Storage Service (Amazon S3)
 - DynamoDB, etc.

[https://github.com/aws/aws-sdk-go](https://github.com/aws/aws-sdk-go)
[https://github.com/aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) (developer preview)
[https://aws.amazon.com/sdk-for-go/](https://aws.amazon.com/sdk-for-go/)

### **Imports**
**Import utility tools and configuration:**

    import "github.com/aws/aws-sdk-go/aws"

**Import service clients, API operations, requests/response structs:**

    import "github.com/aws/aws-sdk-go/service/dynamodb"

### **Configuration**
The goal is to configure the settings for service clients. The only two required options are an AWS Region and credentials. Other optional options are log level and maximum number of retries. These values could be at environment variables or configuration files. Here an example using the configuration set in the code

![](https://lh4.googleusercontent.com/to1qa3r3OQX01fCa55WSUUgmIYJdZeFkU0d1gEkAZEK1Qgg7b47laeQNckI9vvAIDU2JMinCedZGyh7gwLkOtlG1x0FJwEwINSJvYLu3FoxdEAtB1J-vGPBG_BDBfEqW4KvLddzycu4)

**Environment Variables:** By default, the SDK detects AWS credentials set in your environment and uses them to perform requests to AWS

    $ export AWS_REGION=us-west-2
    $ export AWS_ACCESS_KEY_ID=YOUR_AKID
    $ export AWS_SECRET_ACCESS_KEY=YOUR_SECRET_KEY
    $ export AWS_SESSION_TOKEN=YOUR_TOKEN
**Credentials File:** You can use a shared credentials file at .aws/credentials
![](https://lh3.googleusercontent.com/z8grx6oNjL3SawnLZxOhjsQyDEs3wkmZWqHBC6hXHYUPRhK6fCI2mYXKT1T7jx29r03fTH_EU89pq2cX9nCVS6YOuU4BE319c8YAI_K1CeUg6L5ARZk4V-KuGJ4ORny7y7PvrMQLFhU)

### **Example: Upload object to S3**
![](https://lh4.googleusercontent.com/R--CgriayuWm75_Y__lvkacTvMz4Goq6--CzUioiMoM0F0_6qsDFxhFY8Bs-q3x88mAPBnIttBF4RsDKAMrud38RGXR-AZJmtO1Q2-TC4UwTMkZvpWv7VSJJwfU45817CrYsYqjNSRE)

# **DynamoDB API**
## **AWS SDK for DynamoDB**
![](https://lh5.googleusercontent.com/ea84_-2ziYDg7V1SaFnf2ON_LnwRrkvnOBHJUU2kn_e_0mQbwnSTcTsxsDmYyXjESe7YaWeE0UorRACvNxNGJ-2IPvINTq7Ih5-ydnVUiLu7s8NG14cXR3zTCGYUzLglaq1FG910RFE)

## **Handle Tables**

 - **CreateTable:** Creates a new table and one or more secondary indexes
 - **DescribeTable:** Returns information about a table schema
 - **ListTables:** Returns the table names
 - **UpdateTable:** Change the table settings
 - **DeleteTable:** Removes a table

## **CRUD operations**
**Creating Data:** PutItem; BatchWriteItem
**Reading Data:** GetItem; BatchGetItem; Query; Scan
**Updating Data:** UpdateItem
**Deleting Data:** DeleteItem; BatchWriteItem
**… Streams and Transactions:** https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.API.html

## **Define a table**

 - Define table name
 - Define primary key
	 - Partition key
	 - Sort key (optional)
 - Define table attributes
 - Define secondary indexes (optional)
 
 **Topic for another talk:** DynamoDB modeling techniques

## **Data types**

 - **Scalar Types:** Number, string, binary, boolean, and null
 - **Document Types:** Represent a complex structure with nested attributes (JSON). List and map
 - **Set Types:** Represent multiple scalar values. String set, number set, and binary set.

## **Data type descriptors**
Data type descriptors are **tokens** that tell DynamoDB how to interpret each attribute. Here the list:
-   S – String
-   N – Number
-   B – Binary
-   BOOL – Boolean
-   NULL – Null
-   M – Map
-   L – List
-   SS – String Set
-   NS – Number Set
-   BS – Binary Set

# **Best Practices**
https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/best-practices.html
