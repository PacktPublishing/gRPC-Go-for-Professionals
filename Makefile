all:

validate:
	./scripts/startupvalidator.sh -c

update:
	./scripts/updateall.sh

clean:
	./scripts/cleanall.sh