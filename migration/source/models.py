from django.db import models

# Create your models here.
class User(models.Model):
    id = models.UUIDField(primary_key=True)
    email = models.CharField(unique=True, max_length=128, null=True)
    email_verified = models.BooleanField()
    mobilephone = models.CharField(unique=True, max_length=32, null=True)
    mobilephone_verified = models.BooleanField()
    active = models.BooleanField()
    first_name = models.CharField(max_length=128, blank=True, null=True)
    last_name = models.CharField(max_length=128, blank=True, null=True)
    gender = models.CharField(max_length=1, blank=True, null=True)
    birth_place = models.CharField(max_length=128, blank=True, null=True)
    date_of_birth = models.DateField(blank=True, null=True)
    register_provider = models.CharField(max_length=128, blank=True, null=True)  # ex: facebook, google, native
    photo = models.CharField(max_length=4096, blank=True, null=True)
    deleted = models.BooleanField()
    created_at = models.DateTimeField(db_index=True)
    updated_at = models.DateTimeField()
    profession = models.CharField(max_length=256, null=True)
    mobilephone_prefix = models.CharField(max_length=128, blank=True, null=True)
    is_admin = models.BooleanField(default=False)
    is_staff = models.BooleanField(default=False)
    must_update_password = models.BooleanField(default=False)

    class Meta:
        db_table = '"users"'