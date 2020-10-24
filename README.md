# Go language programming: Design patterns (DP)

The repository implements various design patterns in the Go Langauage. Design patterns are typical solutions to commonly occurring problems in software design. They are like pre-made blueprints that you can customize to solve a recurring design problem in your code. The design patterns are namely categorized into three part,
* Creational DP
* structural DP
* behavirol DP

## Creational Design Pattern (CDP)
Creational patterns provide various object creation mechanisms, which increase flexibility and reuse of existing code. I have implemented the below widely used creational design patterns,
* **Singleton CDP:** 
Singleton is a creational design pattern that lets you ensure that a class has only one instance, while providing a global access point to this instance.
* **Builder CDP:** 
Builder is a creational design pattern that lets you construct complex objects step by step. The pattern allows you to produce different types and representations of an object using the same construction code.
* **Factory CDP:** 
Factory Method is a creational design pattern that provides an interface for creating objects in a superclass, but allows subclasses to alter the type of objects that will be created.
* **Abstract Factory CDP:** (Factory of factories)
Abstract Factory is a creational design pattern that lets you produce families of related objects without specifying their concrete classes.

## Structural Design Pattern (SDP)
Structural patterns explain how to assemble objects and classes into larger structures while keeping these structures flexible and efficient. I have implemented the below widely used structural design patterns,
* **Adapter SDP:**
Adapter is a structural design pattern that allows objects with incompatible interfaces to collaborate.
* **Bridge SCP:** *(Under Implementation)*
Bridge is a structural design pattern that lets you split a large class or a set of closely related classes into two separate hierarchies—abstraction and implementation—which can be developed independently of each other.
* **Proxy SCP:** *(Under Implementation)*
Proxy is a structural design pattern that lets you provide a substitute or placeholder for another object. A proxy controls access to the original object, allowing you to perform something either before or after the request gets through to the original object.
* **Flyweight SCP:** *(Under Implementation)*
Flyweight is a structural design pattern that lets you fit more objects into the available amount of RAM by sharing common parts of state between multiple objects instead of keeping all of the data in each object.
* **Facade SCP:** *(Under Implementation)*
Facade is a structural design pattern that provides a simplified interface to a library, a framework, or any other complex set of classes.
* **Decorator SCP:** *(Under Implementation)*
Decorator is a structural design pattern that lets you attach new behaviors to objects by placing these objects inside special wrapper objects that contain the behaviors

## Behavioral Design Patterns (BDP) 
Behavioral design patterns are concerned with algorithms and the assignment of responsibilities between objects.
* **Chain of Responsibility BDP:** *(Under Implementation)*
Chain of Responsibility is a behavioral design pattern that lets you pass requests along a chain of handlers. Upon receiving a request, each handler decides either to process the request or to pass it to the next handler in the chain.
* **Command BDP:** *(Under Implementation)*
Command is a behavioral design pattern that turns a request into a stand-alone object that contains all information about the request. This transformation lets you parameterize methods with different requests, delay or queue a request’s execution, and support undoable operations.
* **Iterator BDP:** *(Under Implementation)*
Iterator is a behavioral design pattern that lets you traverse elements of a collection without exposing its underlying representation (list, stack, tree, etc.).
* **Mediator BDP:** *(Under Implementation)*
Mediator is a behavioral design pattern that lets you reduce chaotic dependencies between objects. The pattern restricts direct communications between the objects and forces them to collaborate only via a mediator object.
* **Memento BDP:** *(Under Implementation)*
Memento  is a behavioral design pattern that lets you save and restore the previous state of an object without revealing the details of its implementation.
* **Observer BDP:** *(Under Implementation)*
Observer is a behavioral design pattern that lets you define a subscription mechanism to notify multiple objects about any events that happen to the object they’re observing.
* **State BDP:** *(Under Implementation)*
State is a behavioral design pattern that lets an object alter its behavior when its internal state changes. It appears as if the object changed its class.
* **Strategy BDP:** *(Under Implementation)*
Strategy is a behavioral design pattern that lets you define a family of algorithms, put each of them into a separate class, and make their objects interchangeable.
* **Template Method BDP:** *(Under Implementation)*
Template Method is a behavioral design pattern that defines the skeleton of an algorithm in the superclass but lets subclasses override specific steps of the algorithm without changing its structure.
* **Visitor BDP:** *(Under Implementation)*
Visitor is a behavioral design pattern that lets you separate algorithms from the objects on which they operate.

### About
@Manoj Pawar

