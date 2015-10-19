// +build darwin
// +build arm arm64

AVPlayer* newPlayer(NSString* url) {
  return [AVPlayer playerWithURL:url];
}

void play(AVPlayer* p) {
  [p play];
}

void pause(AVPlayer* p) {
  [p pause];
}