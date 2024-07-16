.PHONY: run

# Foydalanish: make run folderName=your_folder_name
run:
	@ if [ -z "$(folderName)" ]; then \
		echo "Iltimos, folderName ni kiriting. Foydalanish: make run folderName=your_folder_name"; \
		exit 1; \
	fi
	@ if [ -f "$(folderName)/main.go" ]; then \
		cd $(folderName) && go run main.go; \
	else \
		echo "$(folderName) papkasida main.go topilmadi"; \
		exit 1; \
	fi
