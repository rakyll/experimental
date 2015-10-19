// +build darwin
// +build arm arm64

AVPlayer* newPlayer(char* url) {
  return [AVPlayer playerWithURL:[NSString stringWithCString:url encoding:NSASCIIStringEncoding]];
}

void play(AVPlayer* p) {
  [p play];
}

void pause(AVPlayer* p) {
  [p pause];
}