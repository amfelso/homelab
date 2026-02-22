# Bootstrap a node
.PHONY: bootstrap
bootstrap:
	@NODE_IP=$$(./scripts/get-node-ip.sh ${node}); \
	BOOTSTRAP_SCRIPT="bootstrap/init.sh" && source venv/activate > /dev/null; \
	echo "Bootstrapping node $$NODE_ID at $$NODE_IP using $$BOOTSTRAP_SCRIPT..."; \
	scp -i $$SSH_KEY .env $$ADMIN_USER@$$NODE_IP:/tmp/.env > /dev/null; \
	scp -i $$SSH_KEY "$$BOOTSTRAP_SCRIPT" $$ADMIN_USER@$$NODE_IP:/tmp/bootstrap.sh > /dev/null; \
	ssh -i $$SSH_KEY $$ADMIN_USER@$$NODE_IP "chmod +x /tmp/bootstrap.sh && source /tmp/.env && /tmp/bootstrap.sh && rm -f /tmp/bootstrap.sh"; \
	if [ "$$NODE_IP" = "$$NODE_01" ]; then \
		scp -i $$SSH_KEY $$ADMIN_USER@$$NODE_IP:/tmp/.env .env > /dev/null; \
		scp -i $$SSH_KEY $$ADMIN_USER@$$NODE_IP:/tmp/cluster.yaml $$KUBECONFIG > /dev/null; \
		ssh -i $$SSH_KEY $$ADMIN_USER@$$NODE_IP "rm -f /tmp/.env /tmp/cluster.yaml"; \
	else \
		ssh -i $$SSH_KEY $$ADMIN_USER@$$NODE_IP "rm -f /tmp/.env"; \
	fi; \

# SSH into a node
.PHONY: ssh
ssh:
	@NODE_IP=$$(./scripts/get-node-ip.sh ${node}); \
	source venv/activate > /dev/null; \
	echo "Connecting to node $$NODE_VAR at $$NODE_IP..."; \
	ssh -i "$$SSH_KEY" "$$ADMIN_USER@$$NODE_IP";