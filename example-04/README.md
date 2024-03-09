# Goroutines and Channels Problem Statement

## Problem Description

You are tasked with writing a program that simulates a simple messaging system. The program should consist of two components: a sender and a receiver. The sender component should be responsible for generating messages and sending them to the receiver component. The receiver component should receive the messages and print them to the console.

## Requirements

Your program should meet the following requirements:

- The sender component should generate a fixed number of messages.
- The sender component should send each message to the receiver component using a channel.
- The receiver component should receive the messages from the channel and print them to the console.
- The sender and receiver components should run concurrently using goroutines.
