{
  description = "Reusable Go project template with devenv, Just, and golangci-lint";

  outputs =
    { ... }:
    {
      templates.default = {
        path = ./template;
        description = "Go project with devenv, CGO, Just, Cobra, and automatic lint fixes";
        welcomeText = ''
          Enter the environment with `devenv shell`.
          Run `just init github.com/you/project`, then `just check`.
          See README.md for the available tools and recipes.
        '';
      };
    };
}
