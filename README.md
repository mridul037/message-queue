#### Implementing a simple publish-subscribe pattern in Go


 That includes the necessary Broker, Subscriber, and Channel implementations

Explanation:
Broker Structure: The Broker struct manages the subscribers and provides methods to subscribe, unsubscribe, and publish messages.

Subscriber Structure: Each Subscriber has a channel to receive messages and a channel for unsubscribing.

Publishing Messages: When a message is published to a topic, it is sent to all subscribers of that topic.

Goroutine for Subscriber: The subscriber runs in its own goroutine, listening for messages and handling the unsubscribe signal.

Unsubscribing: When unsubscribing, the broker removes the subscriber from its list and closes its channel.

#### channels
Go, channels are a powerful feature used for communication between goroutines. 
They allow you to send and receive values, providing a way to synchronize and coordinate tasks running concurrently.
 

to run app:
``` go run script.js ```

