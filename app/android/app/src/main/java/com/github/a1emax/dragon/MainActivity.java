package com.github.a1emax.dragon;

import android.os.Bundle;
import android.util.Log;

import androidx.appcompat.app.AppCompatActivity;
import androidx.core.view.WindowCompat;
import androidx.core.view.WindowInsetsCompat;
import androidx.core.view.WindowInsetsControllerCompat;

import com.github.a1emax.dragon.go.ebitenmobileview.Ebitenmobileview;
import com.github.a1emax.dragon.go.intern.EbitenView;
import com.github.a1emax.dragon.go.intern.Intern;

import java.util.Objects;

import go.Seq;

public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        // EdgeToEdge.enable(this);
        setContentView(R.layout.activity_main);
        // ViewCompat.setOnApplyWindowInsetsListener(findViewById(R.id.main), (v, insets) -> {
        //    Insets systemBars = insets.getInsets(WindowInsetsCompat.Type.systemBars());
        //    v.setPadding(systemBars.left, systemBars.top, systemBars.right, systemBars.bottom);
        //    return insets;
        // });

        // Hide system bars.
        WindowInsetsControllerCompat windowInsetsController =
                WindowCompat.getInsetsController(getWindow(), getWindow().getDecorView());
        windowInsetsController.setSystemBarsBehavior(
                WindowInsetsControllerCompat.BEHAVIOR_SHOW_TRANSIENT_BARS_BY_SWIPE
        );
        windowInsetsController.hide(WindowInsetsCompat.Type.systemBars());

        // Get directory with both read and write access.
        java.io.File externalFilesDir = getExternalFilesDir(null);
        String externalFilesDirPath = Objects.requireNonNull(externalFilesDir).getPath();

        try {
            Intern.setFilesDir(externalFilesDirPath);
            Intern.activate();
        } catch (Exception e) {
            logGoError(e);
        }

        Seq.setContext(getApplicationContext());
    }

    // EbitenView.suspendGame and EbitenView.resumeGame should be called in onPause and onResume
    // respectively. However, it sometimes leads to a bug that causes the application to restart
    // when resuming, so for now it's enough to call the corresponding Ebitenviewmobile methods.
    private EbitenView getEbitenView() {
        return (EbitenView) this.findViewById(R.id.ebitenview);
    }

    @Override
    protected void onPause() {
        super.onPause();

        try {
            Intern.suspend();
            Ebitenmobileview.suspend();
        } catch (final Exception e) {
            logGoError(e);
        }
    }

    @Override
    protected void onResume() {
        super.onResume();

        try {
            Ebitenmobileview.resume();
            Intern.resume();
        } catch (final Exception e) {
            logGoError(e);
        }
    }

    private void logGoError(Exception e) {
        Log.e("go", e.toString());
    }
}