#import <AVFoundation/AVPlayer.h>
#import <CoreFoundation/CoreFoundation.h>

AVPlayer* newPlayer(char* url) {
  NSURL* u = [NSURL URLWithString: [NSString stringWithCString:url encoding:NSASCIIStringEncoding]];
  return [AVPlayer playerWithURL: u];
}

void play(AVPlayer* p) {
  [p play];
}
