package v1

// This file contains methods that can be used by the go-restful package to generate Swagger
// documentation for the object types found in 'types.go' This file is automatically generated
// by hack/update-generated-swagger-descriptions.sh and should be run after a full build of OpenShift.
// ==== DO NOT EDIT THIS FILE MANUALLY ====

var map_BinaryBuildRequestOptions = map[string]string{
	"":                        "BinaryBuildRequestOptions are the options required to fully speficy a binary build request",
	"metadata":                "metadata for BinaryBuildRequestOptions.",
	"asFile":                  "asFile determines if the binary should be created as a file within the source rather than extracted as an archive",
	"revision.commit":         "revision.commit is the value identifying a specific commit",
	"revision.message":        "revision.message is the description of a specific commit",
	"revision.authorName":     "revision.authorName of the source control user",
	"revision.authorEmail":    "revision.authorEmail of the source control user",
	"revision.committerName":  "revision.committerName of the source control user",
	"revision.committerEmail": "revision.committerEmail of the source control user",
}

func (BinaryBuildRequestOptions) SwaggerDoc() map[string]string {
	return map_BinaryBuildRequestOptions
}

var map_BinaryBuildSource = map[string]string{
	"":       "BinaryBuildSource describes a binary file to be used for the Docker and Source build strategies, where the file will be extracted and used as the build source.",
	"asFile": "asFile indicates that the provided binary input should be considered a single file within the build input. For example, specifying \"webapp.war\" would place the provided binary as `/webapp.war` for the builder. If left empty, the Docker and Source build strategies assume this file is a zip, tar, or tar.gz file and extract it as the source. The custom strategy receives this binary as standard input. This filename may not contain slashes or be '..' or '.'.",
}

func (BinaryBuildSource) SwaggerDoc() map[string]string {
	return map_BinaryBuildSource
}

var map_Build = map[string]string{
	"":         "Build encapsulates the inputs needed to produce a new deployable image, as well as the status of the execution and a reference to the Pod which executed the build.",
	"metadata": "Standard object's metadata.",
	"spec":     "spec is all the inputs used to execute the build.",
	"status":   "status is the current status of the build.",
}

func (Build) SwaggerDoc() map[string]string {
	return map_Build
}

var map_BuildConfig = map[string]string{
	"":         "Build configurations define a build process for new Docker images. There are three types of builds possible - a Docker build using a Dockerfile, a Source-to-Image build that uses a specially prepared base image that accepts source code that it can make runnable, and a custom build that can run // arbitrary Docker images as a base and accept the build parameters. Builds run on the cluster and on completion are pushed to the Docker registry specified in the \"output\" section. A build can be triggered via a webhook, when the base image changes, or when a user manually requests a new build be // created.\n\nEach build created by a build configuration is numbered and refers back to its parent configuration. Multiple builds can be triggered at once. Builds that do not have \"output\" set can be used to test code or run a verification build.",
	"metadata": "metadata for BuildConfig.",
	"spec":     "spec holds all the input necessary to produce a new build, and the conditions when to trigger them.",
	"status":   "status holds any relevant information about a build config",
}

func (BuildConfig) SwaggerDoc() map[string]string {
	return map_BuildConfig
}

var map_BuildConfigList = map[string]string{
	"":         "BuildConfigList is a collection of BuildConfigs.",
	"metadata": "metadata for BuildConfigList.",
	"items":    "items is a list of build configs",
}

func (BuildConfigList) SwaggerDoc() map[string]string {
	return map_BuildConfigList
}

var map_BuildConfigSpec = map[string]string{
	"":          "BuildConfigSpec describes when and how builds are created",
	"triggers":  "triggers determine how new Builds can be launched from a BuildConfig. If no triggers are defined, a new build can only occur as a result of an explicit client build creation.",
	"runPolicy": "RunPolicy describes how the new build created from this build configuration will be scheduled for execution. This is optional, if not specified we default to \"Serial\".",
}

func (BuildConfigSpec) SwaggerDoc() map[string]string {
	return map_BuildConfigSpec
}

var map_BuildConfigStatus = map[string]string{
	"":            "BuildConfigStatus contains current state of the build config object.",
	"lastVersion": "lastVersion is used to inform about number of last triggered build.",
}

func (BuildConfigStatus) SwaggerDoc() map[string]string {
	return map_BuildConfigStatus
}

var map_BuildList = map[string]string{
	"":         "BuildList is a collection of Builds.",
	"metadata": "metadata for BuildList.",
	"items":    "items is a list of builds",
}

func (BuildList) SwaggerDoc() map[string]string {
	return map_BuildList
}

var map_BuildLog = map[string]string{
	"": "BuildLog is the (unused) resource associated with the build log redirector",
}

func (BuildLog) SwaggerDoc() map[string]string {
	return map_BuildLog
}

var map_BuildLogOptions = map[string]string{
	"":             "BuildLogOptions is the REST options for a build log",
	"container":    "cointainer for which to stream logs. Defaults to only container if there is one container in the pod.",
	"follow":       "follow if true indicates that the build log should be streamed until the build terminates.",
	"previous":     "previous returns previous build logs. Defaults to false.",
	"sinceSeconds": "sinceSeconds is a relative time in seconds before the current time from which to show logs. If this value precedes the time a pod was started, only logs since the pod start will be returned. If this value is in the future, no logs will be returned. Only one of sinceSeconds or sinceTime may be specified.",
	"sinceTime":    "sinceTime is an RFC3339 timestamp from which to show logs. If this value precedes the time a pod was started, only logs since the pod start will be returned. If this value is in the future, no logs will be returned. Only one of sinceSeconds or sinceTime may be specified.",
	"timestamps":   "timestamps, If true, add an RFC3339 or RFC3339Nano timestamp at the beginning of every line of log output. Defaults to false.",
	"tailLines":    "tailLines, If set, is the number of lines from the end of the logs to show. If not specified, logs are shown from the creation of the container or sinceSeconds or sinceTime",
	"limitBytes":   "limitBytes, If set, is the number of bytes to read from the server before terminating the log output. This may not display a complete final line of logging, and may return slightly more or slightly less than the specified limit.",
	"nowait":       "noWait if true causes the call to return immediately even if the build is not available yet. Otherwise the server will wait until the build has started.",
	"version":      "version of the build for which to view logs.",
}

func (BuildLogOptions) SwaggerDoc() map[string]string {
	return map_BuildLogOptions
}

var map_BuildOutput = map[string]string{
	"":            "BuildOutput is input to a build strategy and describes the Docker image that the strategy should produce.",
	"to":          "to defines an optional location to push the output of this build to. Kind must be one of 'ImageStreamTag' or 'DockerImage'. This value will be used to look up a Docker image repository to push to. In the case of an ImageStreamTag, the ImageStreamTag will be looked for in the namespace of the build unless Namespace is specified.",
	"pushSecret":  "PushSecret is the name of a Secret that would be used for setting up the authentication for executing the Docker push to authentication enabled Docker Registry (or Docker Hub).",
	"imageLabels": "imageLabels define a list of labels that are applied to the resulting image. If there are multiple labels with the same name then the last one in the list is used.",
}

func (BuildOutput) SwaggerDoc() map[string]string {
	return map_BuildOutput
}

var map_BuildPostCommitSpec = map[string]string{
	"":        "A BuildPostCommitSpec holds a build post commit hook specification. The hook executes a command in a temporary container running the build output image, immediately after the last layer of the image is committed and before the image is pushed to a registry. The command is executed with the current working directory ($PWD) set to the image's WORKDIR.\n\nThe build will be marked as failed if the hook execution fails. It will fail if the script or command return a non-zero exit code, or if there is any other error related to starting the temporary container.\n\nThere are five different ways to configure the hook. As an example, all forms below are equivalent and will execute `rake test --verbose`.\n\n1. Shell script:\n\n       \"postCommit\": {\n         \"script\": \"rake test --verbose\",\n       }\n\n    The above is a convenient form which is equivalent to:\n\n       \"postCommit\": {\n         \"command\": [\"/bin/sh\", \"-ic\"],\n         \"args\":    [\"rake test --verbose\"]\n       }\n\n2. A command as the image entrypoint:\n\n       \"postCommit\": {\n         \"commit\": [\"rake\", \"test\", \"--verbose\"]\n       }\n\n    Command overrides the image entrypoint in the exec form, as documented in\n    Docker: https://docs.docker.com/engine/reference/builder/#entrypoint.\n\n3. Pass arguments to the default entrypoint:\n\n       \"postCommit\": {\n\t\t      \"args\": [\"rake\", \"test\", \"--verbose\"]\n\t      }\n\n    This form is only useful if the image entrypoint can handle arguments.\n\n4. Shell script with arguments:\n\n       \"postCommit\": {\n         \"script\": \"rake test $1\",\n         \"args\":   [\"--verbose\"]\n       }\n\n    This form is useful if you need to pass arguments that would otherwise be\n    hard to quote properly in the shell script. In the script, $0 will be\n    \"/bin/sh\" and $1, $2, etc, are the positional arguments from Args.\n\n5. Command with arguments:\n\n       \"postCommit\": {\n         \"command\": [\"rake\", \"test\"],\n         \"args\":    [\"--verbose\"]\n       }\n\n    This form is equivalent to appending the arguments to the Command slice.\n\nIt is invalid to provide both Script and Command simultaneously. If none of the fields are specified, the hook is not executed.",
	"command": "command is the command to run. It may not be specified with Script. This might be needed if the image doesn't have `/bin/sh`, or if you do not want to use a shell. In all other cases, using Script might be more convenient.",
	"args":    "args is a list of arguments that are provided to either Command, Script or the Docker image's default entrypoint. The arguments are placed immediately after the command to be run.",
	"script":  "script is a shell script to be run with `/bin/sh -ic`. It may not be specified with Command. Use Script when a shell script is appropriate to execute the post build hook, for example for running unit tests with `rake test`. If you need control over the image entrypoint, or if the image does not have `/bin/sh`, use Command and/or Args. The `-i` flag is needed to support CentOS and RHEL images that use Software Collections (SCL), in order to have the appropriate collections enabled in the shell. E.g., in the Ruby image, this is necessary to make `ruby`, `bundle` and other binaries available in the PATH.",
}

func (BuildPostCommitSpec) SwaggerDoc() map[string]string {
	return map_BuildPostCommitSpec
}

var map_BuildRequest = map[string]string{
	"":                 "BuildRequest is the resource used to pass parameters to build generator",
	"metadata":         "metadata for BuildRequest.",
	"revision":         "revision is the information from the source for a specific repo snapshot.",
	"triggeredByImage": "triggeredByImage is the Image that triggered this build.",
	"from":             "from is the reference to the ImageStreamTag that triggered the build.",
	"binary":           "binary indicates a request to build from a binary provided to the builder",
	"lastVersion":      "lastVersion (optional) is the LastVersion of the BuildConfig that was used to generate the build. If the BuildConfig in the generator doesn't match, a build will not be generated.",
	"env":              "env contains additional environment variables you want to pass into a builder container",
	"triggeredBy":      "triggeredBy describes which triggers started the most recent update to the build configuration and contains information about those triggers.",
}

func (BuildRequest) SwaggerDoc() map[string]string {
	return map_BuildRequest
}

var map_BuildSource = map[string]string{
	"":             "BuildSource is the SCM used for the build.",
	"type":         "type of build input to accept",
	"binary":       "binary builds accept a binary as their input. The binary is generally assumed to be a tar, gzipped tar, or zip file depending on the strategy. For Docker builds, this is the build context and an optional Dockerfile may be specified to override any Dockerfile in the build context. For Source builds, this is assumed to be an archive as described above. For Source and Docker builds, if binary.asFile is set the build will receive a directory with a single file. contextDir may be used when an archive is provided. Custom builds will receive this binary as input on STDIN.",
	"dockerfile":   "dockerfile is the raw contents of a Dockerfile which should be built. When this option is specified, the FROM may be modified based on your strategy base image and additional ENV stanzas from your strategy environment will be added after the FROM, but before the rest of your Dockerfile stanzas. The Dockerfile source type may be used with other options like git - in those cases the Git repo will have any innate Dockerfile replaced in the context dir.",
	"git":          "git contains optional information about git build source",
	"images":       "images describes a set of images to be used to provide source for the build",
	"contextDir":   "contextDir specifies the sub-directory where the source code for the application exists. This allows to have buildable sources in directory other than root of repository.",
	"sourceSecret": "sourceSecret is the name of a Secret that would be used for setting up the authentication for cloning private repository. The secret contains valid credentials for remote repository, where the data's key represent the authentication method to be used and value is the base64 encoded credentials. Supported auth methods are: ssh-privatekey.",
	"secrets":      "secrets represents a list of secrets and their destinations that will be used only for the build.",
}

func (BuildSource) SwaggerDoc() map[string]string {
	return map_BuildSource
}

var map_BuildSpec = map[string]string{
	"":            "BuildSpec has the information to represent a build and also additional information about a build",
	"triggeredBy": "triggeredBy describes which triggers started the most recent update to the build configuration and contains information about those triggers.",
}

func (BuildSpec) SwaggerDoc() map[string]string {
	return map_BuildSpec
}

var map_BuildStatus = map[string]string{
	"":                           "BuildStatus contains the status of a build",
	"phase":                      "phase is the point in the build lifecycle.",
	"cancelled":                  "cancelled describes if a cancel event was triggered for the build.",
	"reason":                     "reason is a brief CamelCase string that describes any failure and is meant for machine parsing and tidy display in the CLI.",
	"message":                    "message is a human-readable message indicating details about why the build has this status.",
	"startTimestamp":             "startTimestamp is a timestamp representing the server time when this Build started running in a Pod. It is represented in RFC3339 form and is in UTC.",
	"completionTimestamp":        "completionTimestamp is a timestamp representing the server time when this Build was finished, whether that build failed or succeeded.  It reflects the time at which the Pod running the Build terminated. It is represented in RFC3339 form and is in UTC.",
	"duration":                   "duration contains time.Duration object describing build time.",
	"outputDockerImageReference": "outputDockerImageReference contains a reference to the Docker image that will be built by this build. Its value is computed from Build.Spec.Output.To, and should include the registry address, so that it can be used to push and pull the image.",
	"config":                     "config is an ObjectReference to the BuildConfig this Build is based on.",
	"output":                     "output describes the Docker image the build has produced.",
}

func (BuildStatus) SwaggerDoc() map[string]string {
	return map_BuildStatus
}

var map_BuildStatusOutput = map[string]string{
	"":   "BuildStatusOutput contains the status of the built image.",
	"to": "to describes the status of the built image being pushed to a registry.",
}

func (BuildStatusOutput) SwaggerDoc() map[string]string {
	return map_BuildStatusOutput
}

var map_BuildStatusOutputTo = map[string]string{
	"":            "BuildStatusOutputTo describes the status of the built image with regards to image registry to which it was supposed to be pushed.",
	"imageDigest": "imageDigest is the digest of the built Docker image. The digest uniquely identifies the image in the registry to which it was pushed.\n\nPlease note that this field may not always be set even if the push completes successfully - e.g. when the registry returns no digest or returns it in a format that the builder doesn't understand.",
}

func (BuildStatusOutputTo) SwaggerDoc() map[string]string {
	return map_BuildStatusOutputTo
}

var map_BuildStrategy = map[string]string{
	"":                        "BuildStrategy contains the details of how to perform a build.",
	"type":                    "type is the kind of build strategy.",
	"dockerStrategy":          "dockerStrategy holds the parameters to the Docker build strategy.",
	"sourceStrategy":          "sourceStrategy holds the parameters to the Source build strategy.",
	"customStrategy":          "customStrategy holds the parameters to the Custom build strategy",
	"jenkinsPipelineStrategy": "JenkinsPipelineStrategy holds the parameters to the Jenkins Pipeline build strategy. This strategy is in tech preview.",
}

func (BuildStrategy) SwaggerDoc() map[string]string {
	return map_BuildStrategy
}

var map_BuildTriggerCause = map[string]string{
	"":                 "BuildTriggerCause holds information about a triggered build. It is used for displaying build trigger data for each build and build configuration in oc describe. It is also used to describe which triggers led to the most recent update in the build configuration.",
	"message":          "message is used to store a human readable message for why the build was triggered. E.g.: \"Manually triggered by user\", \"Configuration change\",etc.",
	"genericWebHook":   "genericWebHook holds data about a builds generic webhook trigger.",
	"githubWebHook":    "gitHubWebHook represents data for a GitHub webhook that fired a specific build.",
	"imageChangeBuild": "imageChangeBuild stores information about an imagechange event that triggered a new build.",
}

func (BuildTriggerCause) SwaggerDoc() map[string]string {
	return map_BuildTriggerCause
}

var map_BuildTriggerPolicy = map[string]string{
	"":            "BuildTriggerPolicy describes a policy for a single trigger that results in a new Build.",
	"type":        "type is the type of build trigger",
	"github":      "github contains the parameters for a GitHub webhook type of trigger",
	"generic":     "generic contains the parameters for a Generic webhook type of trigger",
	"imageChange": "imageChange contains parameters for an ImageChange type of trigger",
}

func (BuildTriggerPolicy) SwaggerDoc() map[string]string {
	return map_BuildTriggerPolicy
}

var map_CommonSpec = map[string]string{
	"":                          "CommonSpec encapsulates all the inputs necessary to represent a build.",
	"serviceAccount":            "serviceAccount is the name of the ServiceAccount to use to run the pod created by this build. The pod will be allowed to use secrets referenced by the ServiceAccount",
	"source":                    "source describes the SCM in use.",
	"revision":                  "revision is the information from the source for a specific repo snapshot. This is optional.",
	"strategy":                  "strategy defines how to perform a build.",
	"output":                    "output describes the Docker image the Strategy should produce.",
	"resources":                 "resources computes resource requirements to execute the build.",
	"postCommit":                "postCommit is a build hook executed after the build output image is committed, before it is pushed to a registry.",
	"completionDeadlineSeconds": "completionDeadlineSeconds is an optional duration in seconds, counted from the time when a build pod gets scheduled in the system, that the build may be active on a node before the system actively tries to terminate the build; value must be positive integer",
	"nodeSelector":              "nodeSelector is a selector which must be true for the build pod to fit on a node If nil, it can be overridden by default build nodeselector values for the cluster. If set to an empty map or a map with any values, default build nodeselector values are ignored.",
}

func (CommonSpec) SwaggerDoc() map[string]string {
	return map_CommonSpec
}

var map_CustomBuildStrategy = map[string]string{
	"":                   "CustomBuildStrategy defines input parameters specific to Custom build.",
	"from":               "from is reference to an DockerImage, ImageStreamTag, or ImageStreamImage from which the docker image should be pulled",
	"pullSecret":         "pullSecret is the name of a Secret that would be used for setting up the authentication for pulling the Docker images from the private Docker registries",
	"env":                "env contains additional environment variables you want to pass into a builder container",
	"exposeDockerSocket": "exposeDockerSocket will allow running Docker commands (and build Docker images) from inside the Docker container.",
	"forcePull":          "forcePull describes if the controller should configure the build pod to always pull the images for the builder or only pull if it is not present locally",
	"secrets":            "secrets is a list of additional secrets that will be included in the build pod",
	"buildAPIVersion":    "buildAPIVersion is the requested API version for the Build object serialized and passed to the custom builder",
}

func (CustomBuildStrategy) SwaggerDoc() map[string]string {
	return map_CustomBuildStrategy
}

var map_DockerBuildStrategy = map[string]string{
	"":               "DockerBuildStrategy defines input parameters specific to Docker build.",
	"from":           "from is reference to an DockerImage, ImageStreamTag, or ImageStreamImage from which the docker image should be pulled the resulting image will be used in the FROM line of the Dockerfile for this build.",
	"pullSecret":     "pullSecret is the name of a Secret that would be used for setting up the authentication for pulling the Docker images from the private Docker registries",
	"noCache":        "noCache if set to true indicates that the docker build must be executed with the --no-cache=true flag",
	"env":            "env contains additional environment variables you want to pass into a builder container",
	"forcePull":      "forcePull describes if the builder should pull the images from registry prior to building.",
	"dockerfilePath": "dockerfilePath is the path of the Dockerfile that will be used to build the Docker image, relative to the root of the context (contextDir).",
}

func (DockerBuildStrategy) SwaggerDoc() map[string]string {
	return map_DockerBuildStrategy
}

var map_GenericWebHookCause = map[string]string{
	"":         "GenericWebHookCause holds information about a generic WebHook that triggered a build.",
	"revision": "revision is an optional field that stores the git source revision information of the generic webhook trigger when it is available.",
	"secret":   "secret is the obfuscated webhook secret that triggered a build.",
}

func (GenericWebHookCause) SwaggerDoc() map[string]string {
	return map_GenericWebHookCause
}

var map_GenericWebHookEvent = map[string]string{
	"":     "GenericWebHookEvent is the payload expected for a generic webhook post",
	"type": "type is the type of source repository",
	"git":  "git is the git information if the Type is BuildSourceGit",
	"env":  "env contains additional environment variables you want to pass into a builder container",
}

func (GenericWebHookEvent) SwaggerDoc() map[string]string {
	return map_GenericWebHookEvent
}

var map_GitBuildSource = map[string]string{
	"":    "GitBuildSource defines the parameters of a Git SCM",
	"uri": "uri points to the source that will be built. The structure of the source will depend on the type of build to run",
	"ref": "ref is the branch/tag/ref to build.",
}

func (GitBuildSource) SwaggerDoc() map[string]string {
	return map_GitBuildSource
}

var map_GitHubWebHookCause = map[string]string{
	"":         "GitHubWebHookCause has information about a GitHub webhook that triggered a build.",
	"revision": "revision is the git revision information of the trigger.",
	"secret":   "secret is the obfuscated webhook secret that triggered a build.",
}

func (GitHubWebHookCause) SwaggerDoc() map[string]string {
	return map_GitHubWebHookCause
}

var map_GitInfo = map[string]string{
	"": "GitInfo is the aggregated git information for a generic webhook post",
}

func (GitInfo) SwaggerDoc() map[string]string {
	return map_GitInfo
}

var map_GitSourceRevision = map[string]string{
	"":          "GitSourceRevision is the commit information from a git source for a build",
	"commit":    "commit is the commit hash identifying a specific commit",
	"author":    "author is the author of a specific commit",
	"committer": "committer is the committer of a specific commit",
	"message":   "message is the description of a specific commit",
}

func (GitSourceRevision) SwaggerDoc() map[string]string {
	return map_GitSourceRevision
}

var map_ImageChangeCause = map[string]string{
	"":        "ImageChangeCause contains information about the image that triggered a build",
	"imageID": "imageID is the ID of the image that triggered a a new build.",
	"fromRef": "fromRef contains detailed information about an image that triggered a build.",
}

func (ImageChangeCause) SwaggerDoc() map[string]string {
	return map_ImageChangeCause
}

var map_ImageChangeTrigger = map[string]string{
	"": "ImageChangeTrigger allows builds to be triggered when an ImageStream changes",
	"lastTriggeredImageID": "lastTriggeredImageID is used internally by the ImageChangeController to save last used image ID for build",
	"from":                 "from is a reference to an ImageStreamTag that will trigger a build when updated It is optional. If no From is specified, the From image from the build strategy will be used. Only one ImageChangeTrigger with an empty From reference is allowed in a build configuration.",
}

func (ImageChangeTrigger) SwaggerDoc() map[string]string {
	return map_ImageChangeTrigger
}

var map_ImageLabel = map[string]string{
	"":      "ImageLabel represents a label applied to the resulting image.",
	"name":  "name defines the name of the label. It must have non-zero length.",
	"value": "value defines the literal value of the label.",
}

func (ImageLabel) SwaggerDoc() map[string]string {
	return map_ImageLabel
}

var map_ImageSource = map[string]string{
	"":           "ImageSource is used to describe build source that will be extracted from an image. A reference of type ImageStreamTag, ImageStreamImage or DockerImage may be used. A pull secret can be specified to pull the image from an external registry or override the default service account secret if pulling from the internal registry. A list of paths to copy from the image and their respective destination within the build directory must be specified in the paths array.",
	"from":       "from is a reference to an ImageStreamTag, ImageStreamImage, or DockerImage to copy source from.",
	"paths":      "paths is a list of source and destination paths to copy from the image.",
	"pullSecret": "pullSecret is a reference to a secret to be used to pull the image from a registry If the image is pulled from the OpenShift registry, this field does not need to be set.",
}

func (ImageSource) SwaggerDoc() map[string]string {
	return map_ImageSource
}

var map_ImageSourcePath = map[string]string{
	"":               "ImageSourcePath describes a path to be copied from a source image and its destination within the build directory.",
	"sourcePath":     "sourcePath is the absolute path of the file or directory inside the image to copy to the build directory.",
	"destinationDir": "destinationDir is the relative directory within the build directory where files copied from the image are placed.",
}

func (ImageSourcePath) SwaggerDoc() map[string]string {
	return map_ImageSourcePath
}

var map_JenkinsPipelineBuildStrategy = map[string]string{
	"":                "JenkinsPipelineBuildStrategy holds parameters specific to a Jenkins Pipeline build. This strategy is in tech preview.",
	"jenkinsfilePath": "JenkinsfilePath is the optional path of the Jenkinsfile that will be used to configure the pipeline relative to the root of the context (contextDir). If both JenkinsfilePath & Jenkinsfile are both not specified, this defaults to Jenkinsfile in the root of the specified contextDir.",
	"jenkinsfile":     "Jenkinsfile defines the optional raw contents of a Jenkinsfile which defines a Jenkins pipeline build.",
}

func (JenkinsPipelineBuildStrategy) SwaggerDoc() map[string]string {
	return map_JenkinsPipelineBuildStrategy
}

var map_ProxyConfig = map[string]string{
	"":           "ProxyConfig defines what proxies to use for an operation",
	"httpProxy":  "httpProxy is a proxy used to reach the git repository over http",
	"httpsProxy": "httpsProxy is a proxy used to reach the git repository over https",
	"noProxy":    "noProxy is the list of domains for which the proxy should not be used",
}

func (ProxyConfig) SwaggerDoc() map[string]string {
	return map_ProxyConfig
}

var map_SecretBuildSource = map[string]string{
	"":               "SecretBuildSource describes a secret and its destination directory that will be used only at the build time. The content of the secret referenced here will be copied into the destination directory instead of mounting.",
	"secret":         "secret is a reference to an existing secret that you want to use in your build.",
	"destinationDir": "destinationDir is the directory where the files from the secret should be available for the build time. For the Source build strategy, these will be injected into a container where the assemble script runs. Later, when the script finishes, all files injected will be truncated to zero length. For the Docker build strategy, these will be copied into the build directory, where the Dockerfile is located, so users can ADD or COPY them during docker build.",
}

func (SecretBuildSource) SwaggerDoc() map[string]string {
	return map_SecretBuildSource
}

var map_SecretSpec = map[string]string{
	"":             "SecretSpec specifies a secret to be included in a build pod and its corresponding mount point",
	"secretSource": "secretSource is a reference to the secret",
	"mountPath":    "mountPath is the path at which to mount the secret",
}

func (SecretSpec) SwaggerDoc() map[string]string {
	return map_SecretSpec
}

var map_SourceBuildStrategy = map[string]string{
	"":                 "SourceBuildStrategy defines input parameters specific to an Source build.",
	"from":             "from is reference to an DockerImage, ImageStreamTag, or ImageStreamImage from which the docker image should be pulled",
	"pullSecret":       "pullSecret is the name of a Secret that would be used for setting up the authentication for pulling the Docker images from the private Docker registries",
	"env":              "env contains additional environment variables you want to pass into a builder container",
	"scripts":          "scripts is the location of Source scripts",
	"incremental":      "incremental flag forces the Source build to do incremental builds if true.",
	"forcePull":        "forcePull describes if the builder should pull the images from registry prior to building.",
	"runtimeImage":     "runtimeImage is an optional image that is used to run an application without unneeded dependencies installed. The building of the application is still done in the builder image but, post build, you can copy the needed artifacts in the runtime image for use. This field and the feature it enables are in tech preview.",
	"runtimeArtifacts": "runtimeArtifacts specifies a list of source/destination pairs that will be copied from the builder to the runtime image. sourcePath can be a file or directory. destinationDir must be a directory. destinationDir can also be empty or equal to \".\", in this case it just refers to the root of WORKDIR. This field and the feature it enables are in tech preview.",
}

func (SourceBuildStrategy) SwaggerDoc() map[string]string {
	return map_SourceBuildStrategy
}

var map_SourceControlUser = map[string]string{
	"":      "SourceControlUser defines the identity of a user of source control",
	"name":  "name of the source control user",
	"email": "email of the source control user",
}

func (SourceControlUser) SwaggerDoc() map[string]string {
	return map_SourceControlUser
}

var map_SourceRevision = map[string]string{
	"":     "SourceRevision is the revision or commit information from the source for the build",
	"type": "type of the build source, may be one of 'Source', 'Dockerfile', 'Binary', or 'Images'",
	"git":  "Git contains information about git-based build source",
}

func (SourceRevision) SwaggerDoc() map[string]string {
	return map_SourceRevision
}

var map_WebHookTrigger = map[string]string{
	"":         "WebHookTrigger is a trigger that gets invoked using a webhook type of post",
	"secret":   "secret used to validate requests.",
	"allowEnv": "allowEnv determines whether the webhook can set environment variables; can only be set to true for GenericWebHook.",
}

func (WebHookTrigger) SwaggerDoc() map[string]string {
	return map_WebHookTrigger
}
