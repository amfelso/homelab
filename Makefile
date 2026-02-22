# SSH into a node
.PHONY: ssh
ssh:
	@NODE_ID=${node}; \
	if [ -z "$$NODE_ID" ]; then \
		echo "Node ID not set. Example: make ssh node=1"; exit 1; \
	fi; \
	source venv/activate > /dev/null && NODE_VAR="NODE_0$$NODE_ID" && NODE_IP=$${!NODE_VAR}; \
	echo "Connecting to node $$NODE_VAR at $$NODE_IP..."; \
	ssh -i "$$SSH_KEY" "$$ADMIN_USER@$$NODE_IP";