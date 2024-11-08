class Melodica < Formula
    desc "Melodica is a console-based audio player built with Go"
    homepage "https://github.com/zombocoder/melodica"
    url "https://github.com/zombocoder/melodica/archive/v0.0.2.tar.gz"
    sha256 "6f66853918049596774ccfa098109aa57d8b852bbfddfc88e5773346375cdb8c"
    license "MIT"
    version "0.0.2"
  
    depends_on "go" => :build
  
    def install
      system "go", "build", *std_go_args(output: bin/"melodica"), "./cmd/melodica"
    end
  
    test do
      assert_match "Usage: melodica <playlist.txt>", shell_output("#{bin}/melodica -h", 2)
    end
  end
  