# Internal Notes

echo $CR_PAT | docker login ghcr.io -u USERNAME --password-stdin

# new login with new container registry url and PAT
echo ${{ secrets.CR_PAT }} | docker login ghcr.io -u $GITHUB_ACTOR --password-stdin
# new container registry urls added
docker pull ghcr.io/github/octoshift:latest
docker build . --tag ghcr.io/github/octoshift:$GITHUB_SHA --cache-from ghcr.io/github/octoshift:latest
docker push ghcr.io/github/octoshift:$GITHUB_SHA

