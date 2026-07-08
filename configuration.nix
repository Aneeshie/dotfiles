{ ... }:

{
  nix.enable = false;

  nixpkgs.config.allowUnfree = true;
  nixpkgs.hostPlatform = "aarch64-darwin";

  system.primaryUser = "nara";
  users.users.nara = {
    home = "/Users/nara";
  };
  system.stateVersion = 6;

  system.defaults = {
    NSGlobalDomain = {
      AppleInterfaceStyle = "Dark";
      KeyRepeat = 2;
      InitialKeyRepeat = 15;
      _HIHideMenuBar = true;
      AppleShowAllExtensions = true;
    };

    dock.autohide = true;

    finder = {
      FXPreferredViewStyle = "Nlsv";
      CreateDesktop = false;
    };

    trackpad.Clicking = true;
  };

  nix-homebrew = {
    enable = true;
    user = "nara";
    enableRosetta = true;
  };

  homebrew = {
    enable = true;

    onActivation = {
      cleanup = "zap";
      autoUpdate = true;
      upgrade = true;
    };

    brews = [
    	"herdr"
    ];

    taps = [
      "nikitabobko/tap"
    ];

    casks = [
      "nikitabobko/tap/aerospace"
      "wezterm"
      "raycast"
      "orbstack"
      "aerospace"
      "spotify"
      "discord"
      "tailscale-app"
    ];
  };

  launchd.user.agents.remap-caps-to-esc = {
    serviceConfig = {
      ProgramArguments = [
        "/usr/bin/hidutil"
        "property"
        "--set"
        "{\"UserKeyMapping\":[{\"HIDKeyboardModifierMappingSrc\":0x700000039,\"HIDKeyboardModifierMappingDst\":0x700000029}]}"
      ];
      RunAtLoad = true;
    };
  };
}
