//
//  MochiViewController.h
//  basic
//
//  Created by Kevin Dang on 3/31/17.
//  Copyright © 2017 Mochi. All rights reserved.
//

#import <UIKit/UIKit.h>
@import Mochi;
@class MochiNode;

@interface MochiViewController : UIViewController
+ (NSPointerArray *)viewControllers;
+ (MochiViewController *)viewControllerWithIdentifier:(NSInteger)identifier;
+ (void)render;
- (id)initWithName:(NSString *)name;
- (void)render;
- (void)update:(MochiNode *)node;
@property (nonatomic, readonly) NSInteger identifier;
@end
