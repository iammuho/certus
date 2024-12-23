Feature: Validate SSH access in AWS Security Group
  As a DevOps engineer
  I want to ensure that port 22 is open only to a specific IP address
  So that the infrastructure is secure from unauthorized access

  Scenario: Verify SSH port access is restricted
    Given an AWS security group named "my-security-group"
    When I check the inbound rules for port 22
    Then port 22 should be open to "203.0.113.25/32"
    And port 22 should not be open to any other IP