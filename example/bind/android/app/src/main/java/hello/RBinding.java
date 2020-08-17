package hello;

import android.app.Activity;
import android.support.annotation.IdRes;
import android.support.annotation.LayoutRes;
import android.util.Log;
import android.view.View;
import android.widget.TextView;

public class RBinding {

    private long age;

    public RBinding(long age) {
        this.age = age;
    }

    public static String getStringFromJava() {
        return "Hello from java !!";
    }

    public static TextView findTextViewById(Activity activity, @IdRes int id) {
        TextView tv = activity.findViewById(id);
        return tv;
    }

//    public static void setText(TextView tv, CharSequence contentCS) {
//        tv.setText(contentCS);
//    }

    public static void setText(TextView tv, String content) {
        tv.setText(content);
    }

    public static void m(int i32) {
        Log.d("Java_World", "m : i32 = " + i32);
    }

    public static void m(long i64) {
        Log.d("Java_World", "m : i64 = " + i64);
    }

    public static void m(String xyz) {
        Log.d("Java_World", "m: " + xyz);
    }

    public static void setContentView(Activity activity, @LayoutRes int layoutResId) {
        activity.setContentView(layoutResId);
    }
}
