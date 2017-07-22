#import <UIKit/UIKit.h>
#import <MatchaBridge/MatchaBridge.h>
@class MatchaNodeRoot;

@interface MatchaViewController : UIViewController // view.Root
+ (NSPointerArray *)viewControllers;
+ (MatchaViewController *)viewControllerWithIdentifier:(NSInteger)identifier;

- (id)initWithGoValue:(MatchaGoValue *)value;
- (void)update:(MatchaNodeRoot *)node;
- (NSArray<MatchaGoValue *> *)call:(NSString *)funcId viewId:(int64_t)viewId args:(NSArray<MatchaGoValue *> *)args;
@property (nonatomic, readonly) NSInteger identifier;
@property (nonatomic, assign) BOOL updating;
@end


void MatchaConfigureChildViewController(UIViewController *vc);
