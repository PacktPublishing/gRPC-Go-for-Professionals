# gRPC Go for Professionals

<a href="https://www.packtpub.com/product/grpc-go-for-professionals/9781837638840?utm_source=github&utm_medium=repository&utm_campaign="><img src="https://content.packt.com/B19664/cover_image_small.jpg" alt="gRPC Go for Professionals" height="256px" align="right"></a>

This is the code repository for [gRPC Go for Professionals](https://www.packtpub.com/product/grpc-go-for-professionals/9781837638840?utm_source=github&utm_medium=repository&utm_campaign=), published by Packt.

**Implement, test, and deploy production-grade microservices**

## What is this book about?
In recent years, the popularity of microservice architecture has surged, bringing forth a new set of requirements. Among these, efficient communication between different services takes center stage, and that’s where gRPC shines. This book will take you through creating gRPC servers and clients in an efficient, secure, and scalable way. However, communication is just one aspect of microservices, so this book goes beyond that to show you how to deploy your application on Kubernetes and configure other tools that are needed for making your application more resilient. With these tools at your disposal, you’ll be ready to get started with using gRPC in a microservice architecture.

In gRPC Go for Professionals, you’ll explore core concepts such as message transmission and the role of Protobuf in serialization and deserialization. Through a step-by-step implementation of a TODO list API, you’ll see the different features of gRPC in action. You’ll then learn different approaches for testing your services and debug your API endpoints. Finally, you’ll get to grips with deploying application services via Docker images and Kubernetes.

## This book covers the following exciting features:
- Understand the different API endpoints that gRPC lets you write
- Discover the essential considerations when writing your Protobuf files
- Compile Protobuf code with protoc, Buf, and Bazel for efficient development
- Gain insights into how advanced gRPC concepts work
- Grasp techniques for unit testing and load testing your API
- Get to grips with deploying your microservices with Docker and Kubernetes
- Discover tools to write secure and efficient gRPC code

If you feel this book is for you, get your [copy](https://www.amazon.com/dp/1837638845) today!

## Instructions and Navigations
All of the code is organized into folders. For example, Chapter02.

The code will look like the following:
```proto
message AddTaskRequest {
     string description = 1;
     google.protobuf.Timestamp due_date = 2;
}
```

**Following is what you need for this book:**
Whether you’re interested in microservices or looking to use gRPC in your product, this book is for you. To fully benefit from its contents, you’ll need a solid grasp of Go programming and using a terminal. If you’re already familiar with gRPC, this book will help you to explore the different concepts and tools in depth.

With the following software and hardware list you can run all code files present in the book (Chapter 1-9).
## Software and Hardware List
| Chapter | Software required | OS required |
| -------- | ------------------------------------ | ----------------------------------- |
| 4-9 | Go 1.20.4 | Windows, Mac OS X, and Linux (Any) |
| 2-9 | Protobuf 23.2 | Windows, Mac OS X, and Linux (Any) |
| 3-9 | gRPC 1.55.0 | Windows, Mac OS X, and Linux (Any) |
| 4-9 | Buf 1.15.1  | Windows, Mac OS X, and Linux (Any) |
| 4-9 | Bazel 6.2.1 | Windows, Mac OS X, and Linux (Any) |

We also provide a PDF file that has color images of the screenshots/diagrams used in this book. [Click here to download it]( https://packt.link/LEms7).

## Outline

* [Chapter 1 - Networking Primer](chapter1)
* [Chapter 2 - Protobuf Primer](chapter2)
* [Chapter 3 - Introduction to gRPC](chapter3)
* [Chapter 4 - Setting up a Project](chapter4)
* [Chapter 5 - Types of gRPC Endpoints](chapter5)
* [Chapter 6 - Designing Effective APIs](chapter6)
* [Chapter 7 - Out-of-the-box features](chapter7)
* [Chapter 8 - More Essential Features](chapter8)
* [Chapter 9 - Production-grade APIs](chapter9)

## Contributing

There main ways in which you can contribute are the following:

- Correcting/Adding documentation (in or out the code).
- Ask questions or provide feedback in the Issues.
- Propose changes in the Pull Requests (note that, in order to stay consistent with the book, PR will only be accepted for future version of the book).

## Related products
* Domain-Driven Design with Golang [[Packt]](https://www.packtpub.com/product/domain-driven-design-with-golang/9781804613450?utm_source=github&utm_medium=repository&utm_campaign=) [[Amazon]](https://www.amazon.com/dp/1804613452)

* Event-Driven Architecture in Golang [[Packt]](https://www.packtpub.com/product/event-driven-architecture-in-golang/9781803238012?utm_source=github&utm_medium=repository&utm_campaign=) [[Amazon]](https://www.amazon.com/dp/1803238011)

## Get to Know the Author
**Clément Jean**
is the CTO of Education for Ethiopia, a start-up focusing on educating K-12 students in Ethiopia. On top of that, he is also an online instructor (on Udemy, Linux Foundation, and others) teaching people about different kinds of technologies. In both his occupations, he deals with technologies such as gRPC and how to apply them to real-life use cases. His overall goal is to empower people through education and technology.
