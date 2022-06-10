import com.github.jengelman.gradle.plugins.shadow.tasks.ShadowJar

plugins {
    `java-library`
    id("com.github.johnrengelman.shadow") version "7.+"
}

group = "com._2ndwatch"

java {
    toolchain {
        languageVersion.set(JavaLanguageVersion.of(17))
    }
}

repositories {
    mavenCentral()
}

dependencies {
    implementation("log4j:log4j:1.2.17")
    implementation("org.slf4j:slf4j-log4j12:1.7.36")
    implementation("org.apache.kafka:kafka-streams:2.6.3")
    implementation("com.sparkjava:spark-core:2.9.3")
    implementation("com.google.protobuf:protobuf-java:3.20.1")
    implementation("com.google.protobuf:protobuf-java-util:3.20.1")
    testImplementation("org.apache.kafka:kafka-streams-test-utils:2.6.3")
    testImplementation("org.hamcrest:hamcrest:2.2")
    testImplementation("org.junit.jupiter:junit-jupiter-api:5.8.1")
    testRuntimeOnly("org.junit.jupiter:junit-jupiter-engine:5.8.1")
}

sourceSets.all {
    java.srcDirs("src/main/java", "src/generated/java")
}

tasks.getByName("build") {
    finalizedBy("shadowJar")
}

tasks.getByName<Jar>("jar") {
    enabled = false
}

tasks.withType<ShadowJar> {
    archiveFileName.set(System.getProperty("rootProjectName") + ".jar")
    manifest {
        attributes["Main-Class"] = "com._2ndwatch.Main"
    }
}

tasks.getByName<Test>("test") {
    useJUnitPlatform()
}