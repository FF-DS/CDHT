# CDHT

[![Build Status](https://app.travis-ci.com/FF-DS/CDHT.svg?branch=main)](https://app.travis-ci.com/FF-DS/CDHT)

![alt text](https://res.cloudinary.com/linktender/image/upload/v1626708215/Chord_route_l74bff.png)

_A modification and overlay network implementation of chord based dht_


As of 2013, the number of connected devices to the internet was approximately around 2 billion, and this year it’s forecasted to rise up to 27.1 billion. This begs the question of if there is a better way of connecting and synchronizing these devices with the existing infrastructure. Currently, scaling and replication are usually handled by adding more computing machines to the pool of resources already available(or provided by the site owner). However, this approach has some obvious disadvantages, for example, the site owner is expected to add more servers which introduce a financial overhead, purchasing these machine is just half of the solution because when we add a new machine to existing infrastructure, it usually does not integrate seamlessly, and may require some manual work for configuration, plus when a machine fails the disruption caused might affect the overall health of the system in unexpected ways.

In this research-based project, we would like to explore and experiment with a different variation of CHORD-based DHT network. DHT has been around for quite some time, and there are various forms of proposed implementations (like CAN network, Pastry, Tapestry, etc... ). In this research-based project, we would like to implement an overlay network based on CHORD-based DHT specification and build a practical application on top of it solely for demonstration purposes. The overlay network delivered on this project will be able to decouple the network from the application, with more tolerance for failure, more resilience to attacks, and with self-correcting and self-optimizing by changing the spacing of the jumps based on some acceptable and tolerable time parameter(s).
