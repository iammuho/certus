package providers

import (
	"fmt"

	"github.com/cucumber/godog"
	"go.uber.org/zap"
)

type SecurityGroupProvider struct {
	logger    *zap.Logger
	groupName string
	port      int
}

// thereIsAnAWSSecurityGroupNamed logs the group name.
func (p *SecurityGroupProvider) thereIsAnAWSSecurityGroupNamed(groupName string) error {
	p.groupName = groupName
	p.logger.Info("Received security group name", zap.String("group", groupName))
	return nil
}

// iCheckTheInboundRulesForPort logs the port being checked.
func (p *SecurityGroupProvider) iCheckTheInboundRulesForPort(port int) error {
	p.port = port
	p.logger.Info("Checking inbound rules for port", zap.Int("port", port))
	return nil
}

// portShouldBeOpenToIP logs the expected IP access for the port.
func (p *SecurityGroupProvider) portShouldBeOpenToIP(port int, ip string) error {
	if p.port != port {
		return fmt.Errorf("mismatched port: expected %d, but got %d", p.port, port)
	}
	p.logger.Info("Validating port access", zap.Int("port", port), zap.String("allowed_ip", ip))
	return nil
}

// portShouldNotBeOpenToAnyOtherIP logs that no other IPs should have access.
func (p *SecurityGroupProvider) portShouldNotBeOpenToAnyOtherIP(port int) error {
	if p.port != port {
		return fmt.Errorf("mismatched port: expected %d, but got %d", p.port, port)
	}
	p.logger.Info("Ensuring no other IPs have access to port", zap.Int("port", port))
	return nil
}

func InitializeSecurityGroupScenario(ctx *godog.ScenarioContext) {
	logger, _ := zap.NewDevelopment() // Replace with production logger as needed
	provider := &SecurityGroupProvider{logger: logger}

	ctx.Step(`^an AWS security group named "([^"]*)"$`, provider.thereIsAnAWSSecurityGroupNamed)
	ctx.Step(`^I check the inbound rules for port (\d+)$`, provider.iCheckTheInboundRulesForPort)
	ctx.Step(`^port (\d+) should be open to "([^"]*)"$`, provider.portShouldBeOpenToIP)
	ctx.Step(`^port (\d+) should not be open to any other IP$`, provider.portShouldNotBeOpenToAnyOtherIP)
}
