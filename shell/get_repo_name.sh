
REPO_URL=$(git remote get-url origin)

PROJECT_FOLDER_NAME=$(echo ${REPO_URL}|rev|cut -d/ -f1|rev|sed 's/.git//g')

echo ${PROJECT_FOLDER_NAME}
