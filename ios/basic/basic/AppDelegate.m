//
//  AppDelegate.m
//  basic
//
//  Created by Kevin Dang on 3/30/17.
//  Copyright © 2017 Matcha. All rights reserved.
//

#import "AppDelegate.h"
#import "MatchaViewController.h"
#import "MatchaObjcBridge.h"
#import "MatchaNode.h"
#import "ViewController.h"
@import Matcha;

@interface AppDelegate ()
@end

@implementation AppDelegate


- (BOOL)application:(UIApplication *)application didFinishLaunchingWithOptions:(NSDictionary *)launchOptions {
    [[MatchaObjcBridge sharedBridge] configure];
    
    // [[NSNotificationCenter defaultCenter] addObserver:self selector:@selector(keyboardHeightDidChange:) name:UIKeyboardDidChangeFrameNotification object:nil];
    // [[NSNotificationCenter defaultCenter] addObserver:self selector:@selector(keyboardDidShow:) name:UIKeyboardDidShowNotification object:nil];
    // [[NSNotificationCenter defaultCenter] addObserver:self selector:@selector(keyboardDidHide:) name:UIKeyboardDidHideNotification object:nil];
    
    MatchaGoValue *rootVC = [[[MatchaGoValue alloc] initWithFunc:@"github.com/overcyn/matcha/examples/settings New"] call:nil args:nil][0];
    
    self.window = [[UIWindow alloc] initWithFrame:[[UIScreen mainScreen] bounds]];
    self.window.rootViewController = [[MatchaViewController alloc] initWithGoValue:rootVC];
    [self.window makeKeyAndVisible];
    return YES;
}

- (void)keyboardHeightDidChange:(NSNotification *)note {
    NSLog(@"%@",note);
}

- (void)keyboardDidShow:(NSNotification *)note {
    NSLog(@"%@", note);
}
- (void)keyboardDidHide:(NSNotification *)note {
    NSLog(@"%@", note);
}

- (void)applicationWillResignActive:(UIApplication *)application {
    // Sent when the application is about to move from active to inactive state. This can occur for certain types of temporary interruptions (such as an incoming phone call or SMS message) or when the user quits the application and it begins the transition to the background state.
    // Use this method to pause ongoing tasks, disable timers, and invalidate graphics rendering callbacks. Games should use this method to pause the game.
}


- (void)applicationDidEnterBackground:(UIApplication *)application {
    // Use this method to release shared resources, save user data, invalidate timers, and store enough application state information to restore your application to its current state in case it is terminated later.
    // If your application supports background execution, this method is called instead of applicationWillTerminate: when the user quits.
}


- (void)applicationWillEnterForeground:(UIApplication *)application {
    // Called as part of the transition from the background to the active state; here you can undo many of the changes made on entering the background.
}


- (void)applicationDidBecomeActive:(UIApplication *)application {
    // Restart any tasks that were paused (or not yet started) while the application was inactive. If the application was previously in the background, optionally refresh the user interface.
}


- (void)applicationWillTerminate:(UIApplication *)application {
    // Called when the application is about to terminate. Save data if appropriate. See also applicationDidEnterBackground:.
}


@end
