package src

const (
	Version    = "0.0.4"
	Author     = "Willian Damasceno"
	BinaryPath = "/usr/local/bin/manual"

	GitHubUsername       = "WillianDamasceno"
	GitHubRepositoryName = "manual"
	GitHubRepositoryId   = GitHubUsername + "/" + GitHubRepositoryName

	RepositoryUrl = "https://github.com/" + GitHubRepositoryId
	ReleaseUrl    = "https://api.github.com/repos/" + GitHubRepositoryId + "/releases"
)

func GetBinaryUrl(tagName string) string {
	return RepositoryUrl + "/releases/download/" + tagName + "/manual"
}
