from django.db import models

class Task(models.Model):
    context = models.CharField(max_length=200)
    completed = models.BooleanField(default=False)

    def __str__(self):
        return self.context
    
    class Meta:
        verbose_name = "Task"
        verbose_name_plural = "Tasks"
    

