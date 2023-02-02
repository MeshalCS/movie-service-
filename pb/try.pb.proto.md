hello here I’m trying to Generate the Go code from my proto file.<br />
first of all, I know this is a simple way because I will not write any code!<br />
because we all know it will generate it automatically.<br />
but the big challenge when you do it for a first time.<br />
you will try download the environment you will need it, to make simple command to generate it<br />
like this command:<br />
protoc --go_out="output_directory" protobuf.proto<br />


so I will not stop , I will Trying more<br />

now I install Homebrew by the official Homebrew website: https://brew.sh/ in my terminal to install the protobuf compiler:<br />
AND I run this command:<br />
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"<br />
After that I run this command:<br />
brew install protobuf<br />

but I found a missing mistake on my environment, I’m trying to fixe it.<br />
